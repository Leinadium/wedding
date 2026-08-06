package client

import (
	"context"
	"errors"
	"net/http"

	"leinadium.dev/wedding/internal/store"
)

type Client struct {
	ctx context.Context

	api  *apiClient
	auth string
}

func New(url string) *Client {
	return &Client{
		ctx:  context.Background(),
		api:  NewAPIClient(url, &http.Client{}),
		auth: "",
	}
}

func (c *Client) SetAuth(auth string) {
	c.auth = auth
}

func (c *Client) Products() ([]store.Product, error) {
	type Res struct {
		Products []store.Product `json:"products"`
	}
	res, err := Get[Res](c.api, c.ctx, "/product", nil)
	if err != nil {
		return nil, err
	}
	return res.Products, nil
}

func (c *Client) Purchases() ([]store.Purchase, error) {
	headers, err := c.authReq()
	if err != nil {
		return nil, err
	}

	type Res struct {
		Purchases []store.Purchase `json:"purchases"`
	}
	res, err := Get[Res](c.api, c.ctx, "/purchase", headers)
	if err != nil {
		return nil, err
	}
	return res.Purchases, nil
}

func (c *Client) Invites() ([]store.Invite, error) {
	headers, err := c.authReq()
	if err != nil {
		return nil, err
	}

	type Res struct {
		Invites []store.Invite `json:"invites"`
	}
	res, err := Get[Res](c.api, c.ctx, "/invite", headers)
	if err != nil {
		return nil, err
	}
	return res.Invites, nil
}

func (c *Client) CreateInvite(invite store.Invite) (store.InviteID, error) {
	headers, err := c.authReq()
	if err != nil {
		return "", err
	}

	type Res struct {
		InviteID store.InviteID `json:"id"`
	}
	res, err := Post[store.Invite, Res](c.api, c.ctx, "/invite", headers, invite)
	if err != nil {
		return "", err
	}
	return res.InviteID, nil
}

func (c *Client) DeleteInvite(inviteID store.InviteID) error {
	headers, err := c.authReq()
	if err != nil {
		return err
	}
	return Delete(c.api, c.ctx, "/invite/"+string(inviteID), headers)
}

func (c *Client) Attendees() ([]store.Attendee, error) {
	headers, err := c.authReq()
	if err != nil {
		return nil, err
	}

	type Res struct {
		Attendees []store.Attendee `json:"attendees"`
	}
	res, err := Get[Res](c.api, c.ctx, "/attendee", headers)
	if err != nil {
		return nil, err
	}
	return res.Attendees, nil
}

func (c *Client) DeleteAttendee(attendeeID store.AttendeeID) error {
	headers, err := c.authReq()
	if err != nil {
		return err
	}
	return Delete(c.api, c.ctx, "/attendee/"+string(attendeeID.String()), headers)
}

func (c *Client) authReq() (map[string]string, error) {
	if c.auth == "" {
		return nil, errors.New("no auth provided")
	}
	return map[string]string{"Authorization": c.auth}, nil
}
