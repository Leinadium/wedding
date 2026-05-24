package v1

import (
	"context"
	"fmt"

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

func (s *Service) NewConfirmations(ctx context.Context, confirmations []models.Confirmation) error {
	return s.store.NewConfirmations(ctx, confirmations)
}

func (s *Service) NewRejection(ctx context.Context, rejection models.Rejection) error {
	return s.store.NewRejection(ctx, rejection)
}
