package payment

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stripe/stripe-go/v85"
	"leinadium.dev/wedding/internal/models"
	"leinadium.dev/wedding/internal/sync"
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
	trigger       sync.Trigger
}

func New(p Params) *Service {
	return &Service{
		client:        stripe.NewClient(p.Key),
		successURL:    fmt.Sprintf("%s/purchase?session_id={CHECKOUT_SESSION_ID}", p.Domain),
		webhookSecret: p.WebhookSecret,
	}
}

func (s *Service) AddSyncTrigger(trigger sync.Trigger) {
	s.trigger = trigger
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
		//PaymentMethodTypes: []*string{stripe.String("card"), stripe.String("pix")},
	}

	session, err := s.client.V1CheckoutSessions.Create(ctx, params)
	if err != nil {
		return Session{}, fmt.Errorf("could not create session: %v", err)
	}

	return Session{URL: session.URL}, nil
}

func (s *Service) Sessions(body []byte, signature string) ([]Session, error) {
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

func (s *Service) Products(ctx context.Context, inactive bool) ([]models.Product, error) {
	products := []models.Product{}

	params := &stripe.ProductListParams{
		Active: stripe.Bool(!inactive),
		Expand: []*string{
			stripe.String("data.default_price"),
		},
	}

	req := s.client.V1Products.List(ctx, params)
	for p, err := range req.All(ctx) {
		if err != nil {
			return nil, fmt.Errorf("could not list products: %v", err)
		}
		if p == nil {
			continue
		}

		var price int64
		if p.DefaultPrice != nil {
			price = int64(p.DefaultPrice.UnitAmount)
		}

		products = append(products, models.Product{
			StripeID:  models.ProductID(p.ID),
			Name:      p.Name,
			ImageURL:  firstOrZero(p.Images),
			PriceBRL:  price,
			Purchased: !p.Active,
		})
	}

	return products, nil
}

func (s *Service) Purchase(ctx context.Context, session Session) (models.Purchase, error) {
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

	if s.trigger != nil {
		go s.trigger.Trigger()
	}

	return purchase, nil
}

type Session struct {
	ID  string
	URL string
}

func firstOrZero(slice []string) string {
	if len(slice) == 0 {
		return ""
	}
	return slice[0]
}
