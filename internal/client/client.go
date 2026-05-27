package client

import (
	"context"
	"errors"
	"net/http"

	"leinadium.dev/wedding/internal/models"
	"leinadium.dev/wedding/internal/request"
)

type CLIClient interface {
	SetAuth(auth string)

	Products() ([]models.Product, error)
	Purchases() ([]models.Purchase, error)
	Invites() ([]models.Invite, error)
	CreateInvite(models.Invite) (models.InviteID, error)
	DeleteInvite(models.InviteID) error

	Attendees() ([]models.Attendee, error)
	DeleteAttendee(models.AttendeeID) error
}

type client struct {
	ctx context.Context

	api  *request.APIClient
	auth string
}

func New(url string) CLIClient {
	return &client{
		ctx:  context.Background(),
		api:  request.New(url, &http.Client{}),
		auth: "",
	}
}

func (c *client) SetAuth(auth string) {
	c.auth = auth
}

func (c *client) Products() ([]models.Product, error) {
	type Res struct {
		Products []models.Product `json:"products"`
	}
	res, err := request.Get[Res](c.api, c.ctx, "/product", nil)
	if err != nil {
		return nil, err
	}
	return res.Products, nil
}

func (c *client) Purchases() ([]models.Purchase, error) {
	headers, err := c.authReq()
	if err != nil {
		return nil, err
	}

	type Res struct {
		Purchases []models.Purchase `json:"purchases"`
	}
	res, err := request.Get[Res](c.api, c.ctx, "/purchase", headers)
	if err != nil {
		return nil, err
	}
	return res.Purchases, nil
}

func (c *client) Invites() ([]models.Invite, error) {
	headers, err := c.authReq()
	if err != nil {
		return nil, err
	}

	type Res struct {
		Invites []models.Invite `json:"invites"`
	}
	res, err := request.Get[Res](c.api, c.ctx, "/invite", headers)
	if err != nil {
		return nil, err
	}
	return res.Invites, nil
}

func (c *client) CreateInvite(invite models.Invite) (models.InviteID, error) {
	headers, err := c.authReq()
	if err != nil {
		return "", err
	}

	type Res struct {
		InviteID models.InviteID `json:"id"`
	}
	res, err := request.Post[models.Invite, Res](c.api, c.ctx, "/invite", headers, invite)
	if err != nil {
		return "", err
	}
	return res.InviteID, nil
}

func (c *client) DeleteInvite(inviteID models.InviteID) error {
	headers, err := c.authReq()
	if err != nil {
		return err
	}
	return request.Delete(c.api, c.ctx, "/invite/"+string(inviteID), headers)
}

func (c *client) Attendees() ([]models.Attendee, error) {
	headers, err := c.authReq()
	if err != nil {
		return nil, err
	}

	type Res struct {
		Attendees []models.Attendee `json:"attendees"`
	}
	res, err := request.Get[Res](c.api, c.ctx, "/attendee", headers)
	if err != nil {
		return nil, err
	}
	return res.Attendees, nil
}

func (c *client) DeleteAttendee(attendeeID models.AttendeeID) error {
	headers, err := c.authReq()
	if err != nil {
		return err
	}
	return request.Delete(c.api, c.ctx, "/attendee/"+string(attendeeID.String()), headers)
}

func (c *client) authReq() (map[string]string, error) {
	if c.auth == "" {
		return nil, errors.New("no auth provided")
	}
	return map[string]string{"Authorization": c.auth}, nil
}
