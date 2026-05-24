package server

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	api.POST("/confirmation", server.postConfirmations)
	api.POST("/rejection", server.postRejection)

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

func (s *Server) postConfirmations(c *gin.Context) {
	var confirmations []models.Confirmation
	if err := c.BindJSON(&confirmations); err != nil {
		s.error(c, http.StatusBadRequest, err)
		return
	}

	for i, _ := range confirmations {
		confirmations[i].CreatedAt = time.Now()
	}

	if err := s.svc.NewConfirmations(c.Request.Context(), confirmations); err != nil {
		s.error(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func (s *Server) postRejection(c *gin.Context) {
	var rejection models.Rejection
	if err := c.BindJSON(&rejection); err != nil {
		s.error(c, http.StatusBadRequest, err)
		return
	}

	if err := s.svc.NewRejection(c.Request.Context(), rejection); err != nil {
		s.error(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
