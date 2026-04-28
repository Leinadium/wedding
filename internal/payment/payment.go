package payment

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stripe/stripe-go/v85"
	"leinadium.dev/wedding/internal/models"
)

var (
	currency = "BRL"
)

type Params struct {
	Key           string
	WebhookSecret string
	Domain        string
}

type Service struct {
	client        *stripe.Client
	successURL    string
	webhookSecret string
}

func New(p Params) *Service {
	return &Service{
		client:        stripe.NewClient(p.Key),
		successURL:    fmt.Sprintf("%s/success?session_id={CHECKOUT_SESSION_ID}", p.Domain),
		webhookSecret: p.WebhookSecret,
	}
}

func (s *Service) CreateSession(ctx context.Context, product models.Product) (Session, error) {
	// getting payment link
	params := &stripe.CheckoutSessionCreateParams{
		LineItems: []*stripe.CheckoutSessionCreateLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionCreateLineItemPriceDataParams{
					Currency:   &currency,
					Product:    stripe.String(product.StripeID),
					UnitAmount: stripe.Int64(product.PriceBRL),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(s.successURL),
		// PaymentMethodTypes: {}
	}

	session, err := s.client.V1CheckoutSessions.Create(ctx, params)
	if err != nil {
		return Session{}, fmt.Errorf("could not create session: %v", err)
	}

	return Session{URL: session.URL}, nil
}

func (s *Service) GetSessions(body []byte, signature string) ([]Session, error) {
	// Pass the request body and Stripe-Signature header to ConstructEvent, along with the webhook signing key
	// Use the secret provided by Stripe CLI for local testing
	// or your webhook endpoint's secret.

	// body <- req body
	// header <- req.Header.Get("Stripe-Signature")

	event, err := s.client.ConstructEvent(body, signature, s.webhookSecret)

	if err != nil {
		return nil, fmt.Errorf("could not create event: %v", err)
	}

	var sessions []Session

	if event.Type == stripe.EventTypeCheckoutSessionCompleted ||
		event.Type == stripe.EventTypeCheckoutSessionAsyncPaymentSucceeded {
		var cs stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &cs)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal checkout session: %v", err)
		}
		sessions = append(sessions, Session{URL: cs.URL})
	}
	return sessions, nil
}

func (s *Service) GetPurchase(ctx context.Context, session Session) (models.Purchase, error) {
	// TODO: Make this function safe to run multiple times,
	// even concurrently, with the same session ID

	// TODO: Make sure fulfillment hasn't already been
	// performed for this Checkout Session

	// Retrieve the Checkout Session from the API with line_items expanded
	params := &stripe.CheckoutSessionRetrieveParams{}
	params.AddExpand("line_items")

	cs, _ := s.client.V1CheckoutSessions.Retrieve(ctx, session.ID, params)

	// Check the Checkout Session's payment_status property
	// to determine if fulfillment should be performed
	var purchase models.Purchase

	if cs.PaymentStatus != stripe.CheckoutSessionPaymentStatusUnpaid {
		if cs.LineItems != nil {
			for _, line := range cs.LineItems.Data {
				if line.Price != nil && line.Price.Product != nil {
					purchase.ProductID = line.Price.Product.ID
					purchase.ProductName = line.Price.Product.Name
					purchase.Email = cs.CustomerEmail
					purchase.Price = line.AmountTotal
					purchase.ID = line.ID
				}
			}
		}
	}
	return purchase, nil
}

type Session struct {
	ID  string
	URL string
}
