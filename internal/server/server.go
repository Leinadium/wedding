package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"leinadium.dev/wedding/internal/models"
	v1 "leinadium.dev/wedding/internal/v1"
)

type Params struct {
	AuthSecret string
	StaticDir  string
}

type Server struct {
	engine *gin.Engine

	svc        *v1.Service
	authSecret string
}

func New(svc *v1.Service, p Params) *Server {
	engine := gin.Default()
	_ = engine.SetTrustedProxies(nil)

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	server := &Server{
		svc:        svc,
		engine:     engine,
		authSecret: p.AuthSecret,
	}

	api := engine.Group("/v1")
	api.GET("/product", server.getProducts)
	api.GET("/product/:id/payment", server.getProductPayment)
	api.GET("/purchase", server.getPurchases)
	api.POST("/purchase", server.postPurchase)
	api.POST("/invite", server.postInvite)
	api.GET("/invite/:id", server.getInvite)
	api.PATCH("/invite/:id", server.patchInvite)
	api.DELETE("/invite/:id", server.deleteInvite)
	api.GET("/attendee", server.getAttendees)
	api.PATCH("/attendee/:id", server.patchAttendee)
	api.DELETE("/attendee/:id", server.deleteAttendee)

	if p.StaticDir != "" {
		engine.Static("/", p.StaticDir)
	}

	return server
}

func (s *Server) Run(port int) error {
	return s.engine.Run(fmt.Sprintf(":%d", port))
}

func (s *Server) error(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{"error": err.Error()})
}

func (s *Server) checkAuth(c *gin.Context) bool {
	if s.authSecret == "" {
		return true
	}
	if c.GetHeader("Authorization") != s.authSecret {
		c.AbortWithStatus(http.StatusUnauthorized)
		return false
	}
	return true
}

func (s *Server) getProducts(c *gin.Context) {
	products, err := s.svc.Products(c.Request.Context())
	if err != nil {
		s.error(c, http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (s *Server) getProductPayment(c *gin.Context) {
	pid := c.Param("id")
	payment, err := s.svc.Payment(c.Request.Context(), models.ProductID(pid))
	if err != nil {
		s.error(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"payment": payment})

}

func (s *Server) getPurchases(c *gin.Context) {
	if !s.checkAuth(c) {
		return
	}

	purchases, err := s.svc.Purchases(c.Request.Context())
	if err != nil {
		s.error(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"purchases": purchases})
}

func (s *Server) postPurchase(c *gin.Context) {
	// body <- req body
	// header <- req.Header.Get("Stripe-Signature")

	signature := c.GetHeader("Stripe-Signature")
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		s.error(c, http.StatusInternalServerError, err)
	}
	defer c.Request.Body.Close()

	if err := s.svc.NewPurchase(c.Request.Context(), body, signature); err != nil {
		s.error(c, http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func (s *Server) postInvite(c *gin.Context) {
	if !s.checkAuth(c) {
		return
	}

	var invite models.Invite
	if err := c.BindJSON(&invite); err != nil {
		s.error(c, http.StatusBadRequest, err)
		return
	}

	id, err := s.svc.NewInvite(c.Request.Context(), invite)
	if err != nil {
		s.error(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (s *Server) getInvite(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		s.error(c, http.StatusBadRequest, fmt.Errorf("id is required"))
		return
	}

	invite, err := s.svc.Invite(c.Request.Context(), models.InviteID(id))
	if err != nil {
		s.error(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, invite)
}

func (s *Server) patchInvite(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		s.error(c, http.StatusBadRequest, fmt.Errorf("id is required"))
		return
	}

	var req struct {
		Note string `json:"note"`
	}
	if err := c.BindJSON(&req); err != nil {
		s.error(c, http.StatusBadRequest, err)
		return
	}

	if err := s.svc.UpdateInviteNote(c.Request.Context(), models.InviteID(id), req.Note); err != nil {
		s.error(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (s *Server) deleteInvite(c *gin.Context) {
	if !s.checkAuth(c) {
		return
	}

	id := c.Param("id")
	if id == "" {
		s.error(c, http.StatusBadRequest, fmt.Errorf("id is required"))
		return
	}

	if err := s.svc.DeleteInvite(c.Request.Context(), models.InviteID(id)); err != nil {
		s.error(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (s *Server) getAttendees(c *gin.Context) {
	if !s.checkAuth(c) {
		return
	}

	attendees, err := s.svc.Attendees(c.Request.Context())
	if err != nil {
		s.error(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"attendees": attendees})
}

func (s *Server) patchAttendee(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		s.error(c, http.StatusBadRequest, fmt.Errorf("id is required"))
		return
	}

	var req struct {
		IsChild   bool  `json:"isChild"`
		Confirmed *bool `json:"confirmed"`
	}

	if err := c.BindJSON(&req); err != nil {
		s.error(c, http.StatusBadRequest, err)
		return
	}

	if err := s.svc.UpsertAttendee(c.Request.Context(), uuid.MustParse(id), req.IsChild, req.Confirmed); err != nil {
		s.error(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (s *Server) deleteAttendee(c *gin.Context) {
	if !s.checkAuth(c) {
		return
	}

	id := c.Param("id")
	if id == "" {
		s.error(c, http.StatusBadRequest, fmt.Errorf("id is required"))
		return
	}

	if err := s.svc.DeleteAttendee(c.Request.Context(), uuid.MustParse(id)); err != nil {
		s.error(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
