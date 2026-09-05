package v1

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/guregu/null/v6"
	"leinadium.dev/wedding/internal/notification"
	"leinadium.dev/wedding/internal/payment"
	"leinadium.dev/wedding/internal/store"
)

var (
	ErrNotFound = store.ErrNotFound
)

type Service struct {
	store       store.Store
	payment     *payment.Service
	notificator notification.Notificator
}

func New(
	store store.Store,
	payment *payment.Service,
	notificator notification.Notificator,
) *Service {
	return &Service{store: store, payment: payment, notificator: notificator}
}

func (s *Service) Products(ctx context.Context) ([]PurchasableProduct, error) {
	products, err := s.store.Products(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not obtain products: %v", err)
	}

	purchases, err := s.store.Purchases(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not obtain products: %v", err)
	}
	productsPurchased := make(map[store.ProductID]bool)
	for _, p := range purchases {
		productsPurchased[p.ProductID] = true
	}

	var purchasableProducts []PurchasableProduct
	for _, p := range products {
		purchasableProducts = append(purchasableProducts, PurchasableProduct{
			Product:   p,
			Purchased: productsPurchased[p.StripeID],
		})
	}

	return purchasableProducts, nil
}

func (s *Service) Purchase(ctx context.Context, sessionID string) (*store.Purchase, error) {
	return s.payment.Purchase(ctx, payment.Session{ID: sessionID})
}

func (s *Service) Payment(ctx context.Context, pid store.ProductID) (Payment, error) {
	product, err := s.store.Product(ctx, pid)
	if err != nil {
		return Payment{}, fmt.Errorf("could not get product: %v", err)
	}

	session, err := s.payment.CreateSession(ctx, product)
	if err != nil {
		return Payment{}, fmt.Errorf("could not create payment link: %v", err)
	}
	return Payment{URL: session.URL}, nil
}

func (s *Service) Purchases(ctx context.Context) ([]store.Purchase, error) {
	purchases, err := s.store.Purchases(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get purchases: %v", err)
	}
	return purchases, nil
}

func (s *Service) NewPurchase(ctx context.Context, body []byte, signature string) (bool, error) {
	// body <- req body
	// header <- req.Header.Get("Stripe-Signature")
	//
	// flux:
	// get session from signature header
	// get purchase from session
	// store purchase
	// notify purchase
	session, err := s.payment.Session(body, signature)
	if err != nil {
		return false, fmt.Errorf("could not get session: %v", err)
	}

	if session == nil {
		return false, nil
	}

	purchase, err := s.payment.Purchase(ctx, *session)
	if err != nil {
		return false, fmt.Errorf("could not get purchase: %v", err)
	}

	if purchase == nil {
		return false, errors.New("purchase is nil")
	}

	go func() {
		if s.notificator != nil {
			msg := fmt.Sprintf("new purchase: %s bought %s (%f)", purchase.Name, purchase.ProductName, float64(purchase.Price)/100)
			if err := s.notificator.Notify(ctx, msg); err != nil {
				fmt.Printf("could not notify: %v\n", err)
			}
		}
	}()

	fmt.Printf("storing a purchase: %v\n", purchase.ID)

	return true, s.store.NewPurchase(ctx, *purchase)
}

func (s *Service) NewInvite(ctx context.Context, invite store.Invite) (store.InviteID, error) {
	inviteID, err := s.store.NewInvite(ctx, invite)
	if err != nil {
		return "", fmt.Errorf("could not create invite: %v", err)
	}
	return inviteID, nil
}

func (s *Service) Invite(ctx context.Context, inviteID store.InviteID) (store.Invite, error) {
	invite, err := s.store.Invite(ctx, inviteID)
	if err != nil {
		return store.Invite{}, fmt.Errorf("could not get invite: %v", err)
	}
	return invite, nil
}

func (s *Service) Invites(ctx context.Context) ([]store.Invite, error) {
	invites, err := s.store.Invites(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get invites: %v", err)
	}
	return invites, nil
}

func (s *Service) DeleteInvite(ctx context.Context, inviteID store.InviteID) error {
	if err := s.store.DeleteInvite(ctx, inviteID); err != nil {
		return fmt.Errorf("could not delete invite: %v", err)
	}
	return nil
}

func (s *Service) UpdateInviteNote(ctx context.Context, inviteID store.InviteID, note string) error {
	if err := s.store.UpsertNoteInvite(ctx, inviteID, note); err != nil {
		return fmt.Errorf("could not update invite note: %v", err)
	}
	return nil
}

func (s *Service) Attendees(ctx context.Context) ([]store.Attendee, error) {
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
	attendee.Confirmed = null.BoolFromPtr(confirmed)

	if err := s.store.UpsertAttendee(ctx, attendee); err != nil {
		return fmt.Errorf("could not upsert attendee: %v", err)
	}

	if s.notificator != nil {
		msg := createAttendeeNotification(&attendee)
		if err := s.notificator.Notify(ctx, msg); err != nil {
			return fmt.Errorf("could not notify: %v", err)
		}
	}

	return nil
}

func (s *Service) DeleteAttendee(ctx context.Context, attendeeID uuid.UUID) error {
	if err := s.store.DeleteAttendee(ctx, attendeeID); err != nil {
		return fmt.Errorf("could not delete attendee: %v", err)
	}
	return nil
}

func createAttendeeNotification(a *store.Attendee) string {
	var sb strings.Builder
	sb.WriteString("attendee updated: ")
	sb.WriteString(a.Name)
	if a.IsChild {
		sb.WriteString(" (child)")
	}
	sb.WriteString(" -> ")
	if a.Confirmed.Valid && a.Confirmed.Bool {
		sb.WriteString("confirmed")
	} else if a.Confirmed.Valid {
		sb.WriteString("won't go")
	} else {
		sb.WriteString("not confirmed")
	}
	return sb.String()
}
