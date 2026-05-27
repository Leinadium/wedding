package v1

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"leinadium.dev/wedding/internal/models"
	"leinadium.dev/wedding/internal/payment"
	"leinadium.dev/wedding/internal/store"
)

type Params struct {
}

type Service struct {
	store   store.Service
	payment *payment.Service
}

func New(store store.Service, payment *payment.Service, params Params) *Service {
	return &Service{
		store:   store,
		payment: payment,
	}
}

func (s *Service) Products(ctx context.Context) ([]models.Product, error) {
	products, err := s.store.Products(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not obtain products: %v", err)
	}

	return products, nil
}

func (s *Service) Payment(ctx context.Context, pid models.ProductID) (models.Payment, error) {
	product, err := s.store.Product(ctx, pid)
	if err != nil {
		return models.Payment{}, fmt.Errorf("could not get product: %v", err)
	}

	session, err := s.payment.CreateSession(ctx, product)
	if err != nil {
		return models.Payment{}, fmt.Errorf("could not create payment link: %v", err)
	}
	return models.Payment{URL: session.URL}, nil
}

func (s *Service) Purchases(ctx context.Context) ([]models.Purchase, error) {
	purchases, err := s.store.Purchases(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get purchases: %v", err)
	}
	return purchases, nil
}

func (s *Service) NewPurchase(ctx context.Context, body []byte, signature string) error {
	// body <- req body
	// header <- req.Header.Get("Stripe-Signature")
	//
	// flux:
	// get session from signature header
	// get purchase from session
	// store purchase
	// notify purchase
	sessions, err := s.payment.Sessions(body, signature)
	if err != nil {
		return fmt.Errorf("could not get session: %v", err)
	}

	if len(sessions) != 1 {
		return fmt.Errorf("expected 1 session, got %d", len(sessions))
	}

	session := sessions[0]
	purchase, err := s.payment.Purchase(ctx, session)
	if err != nil {
		return fmt.Errorf("could not get purchase: %v", err)
	}

	return s.store.NewPurchase(ctx, purchase)
}

func (s *Service) NewInvite(ctx context.Context, invite models.Invite) (models.InviteID, error) {
	inviteID, err := s.store.NewInvite(ctx, invite)
	if err != nil {
		return "", fmt.Errorf("could not create invite: %v", err)
	}
	return inviteID, nil
}

func (s *Service) Invite(ctx context.Context, inviteID models.InviteID) (models.Invite, error) {
	invite, err := s.store.Invite(ctx, inviteID)
	if err != nil {
		return models.Invite{}, fmt.Errorf("could not get invite: %v", err)
	}
	return invite, nil
}

func (s *Service) DeleteInvite(ctx context.Context, inviteID models.InviteID) error {
	if err := s.store.DeleteInvite(ctx, inviteID); err != nil {
		return fmt.Errorf("could not delete invite: %v", err)
	}
	return nil
}

func (s *Service) UpdateInviteNote(ctx context.Context, inviteID models.InviteID, note string) error {
	if err := s.store.UpsertNoteInvite(ctx, inviteID, note); err != nil {
		return fmt.Errorf("could not update invite note: %v", err)
	}
	return nil
}

func (s *Service) Attendees(ctx context.Context) ([]models.Attendee, error) {
	attendees, err := s.store.Attendees(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get attendees: %v", err)
	}
	return attendees, nil
}

func (s *Service) UpsertAttendee(ctx context.Context, attendeeID uuid.UUID, isChild bool, confirmed *bool) error {
	attendee, err := s.store.Attendee(ctx, attendeeID)
	if err != nil {
		return fmt.Errorf("could not get attendee: %v", err)
	}

	attendee.IsChild = isChild
	attendee.Confirmed = sql.NullBool{Bool: confirmed != nil && *confirmed, Valid: confirmed != nil}

	if err := s.store.UpsertAttendee(ctx, attendee); err != nil {
		return fmt.Errorf("could not upsert attendee: %v", err)
	}
	return nil
}

func (s *Service) DeleteAttendee(ctx context.Context, attendeeID uuid.UUID) error {
	if err := s.store.DeleteAttendee(ctx, attendeeID); err != nil {
		return fmt.Errorf("could not delete attendee: %v", err)
	}
	return nil
}
