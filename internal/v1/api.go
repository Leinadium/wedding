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

func (s *Service) CreateGuest(ctx context.Context, guest models.Guest) error {
	if err := s.store.CreateGuest(ctx, guest); err != nil {
		return fmt.Errorf("could not create guest: %v", err)
	}
	return nil
}

func (s *Service) GetProducts(ctx context.Context) ([]models.Product, error) {
	products, err := s.store.GetProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not obtain products: %v", err)
	}

	return products, nil
}

func (s *Service) GetPayment(ctx context.Context, pid models.ProductID) (models.Payment, error) {
	product, err := s.store.GetProduct(ctx, pid)
	if err != nil {
		return models.Payment{}, fmt.Errorf("could not get product: %v", err)
	}

	session, err := s.payment.CreateSession(ctx, product)
	if err != nil {
		return models.Payment{}, fmt.Errorf("could not create payment link: %v", err)
	}
	return models.Payment{URL: session.URL}, nil
}

func (s *Service) GetPurchases(ctx context.Context) ([]models.Purchase, error) {
	purchases, err := s.store.GetPurchases(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get purchases: %v", err)
	}
	return purchases, nil
}

func (s *Service) CreatePurchase(ctx context.Context, body []byte, signature string) error {
	// body <- req body
	// header <- req.Header.Get("Stripe-Signature")
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

	return s.store.CreatePurchase(ctx, purchase)
}
