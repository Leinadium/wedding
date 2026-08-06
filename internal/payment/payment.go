package payment

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/stripe/stripe-go/v85"
	"leinadium.dev/wedding/internal/store"
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
		successURL:    createSuccessURL(p.Domain),
		webhookSecret: p.WebhookSecret,
	}
}

func (s *Service) AddSyncTrigger(trigger sync.Trigger) {
	s.trigger = trigger
}

func (s *Service) CreateSession(ctx context.Context, product store.Product) (Session, error) {
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
		CustomerCreation: stripe.String(string(stripe.CheckoutSessionCustomerCreationAlways)),
		NameCollection: &stripe.CheckoutSessionCreateNameCollectionParams{
			Individual: &stripe.CheckoutSessionCreateNameCollectionIndividualParams{
				Enabled:  stripe.Bool(true),
				Optional: stripe.Bool(false),
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

func (s *Service) Session(body []byte, signature string) (*Session, error) {
	// Pass the request body and Stripe-Signature header to ConstructEvent, along with the webhook signing key
	// Use the secret provided by Stripe CLI for local testing
	// or your webhook endpoint's secret.

	// body <- req body
	// header <- req.Header.Get("Stripe-Signature")

	event, err := s.client.ConstructEvent(body, signature, s.webhookSecret)

	if err != nil {
		return nil, fmt.Errorf("could not create event: %v", err)
	}

	if event.Type == stripe.EventTypeCheckoutSessionCompleted ||
		event.Type == stripe.EventTypeCheckoutSessionAsyncPaymentSucceeded {
		var cs stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &cs)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal checkout session: %v", err)
		}
		return &Session{ID: cs.ID, URL: cs.URL}, nil
	}
	return nil, nil
}

func (s *Service) Products(ctx context.Context, inactive bool) ([]store.Product, error) {
	products := []store.Product{}

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

		products = append(products, store.Product{
			StripeID: store.ProductID(p.ID),
			Name:     p.Name,
			ImageURL: firstOrZero(p.Images),
			PriceBRL: price,
			Active:   p.Active,
		})
	}

	return products, nil
}

func (s *Service) Purchase(ctx context.Context, session Session) (*store.Purchase, error) {
	// TODO: Make this function safe to run multiple times,
	// even concurrently, with the same session ID

	// TODO: Make sure fulfillment hasn't already been
	// performed for this Checkout Session

	// Retrieve the Checkout Session from the API with line_items expanded
	params := &stripe.CheckoutSessionRetrieveParams{}
	params.AddExpand("line_items")
	params.AddExpand("customer")

	cs, _ := s.client.V1CheckoutSessions.Retrieve(ctx, session.ID, params)

	// Check the Checkout Session's payment_status property
	// to determine if fulfillment should be performed
	var purchase *store.Purchase

	if cs.PaymentStatus != stripe.CheckoutSessionPaymentStatusUnpaid {
		if cs.LineItems != nil {
			var err error
			purchase, err = stripeIntoProduct(cs.LineItems.Data, cs.Customer)
			if err != nil {
				return nil, err
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

func createSuccessURL(domain string) string {
	return fmt.Sprintf("%s?callback={CHECKOUT_SESSION_ID}", domain)
}

func stripeIntoProduct(lines []*stripe.LineItem, customer *stripe.Customer) (*store.Purchase, error) {
	var purchase store.Purchase
	for _, line := range lines {
		if line.Price != nil && line.Price.Product != nil {
			purchase.ProductID = store.ProductID(line.Price.Product.ID)
			purchase.ProductName = line.Description // "defaults to product name when not set"
			purchase.Price = line.AmountTotal
			purchase.ID = line.ID
			purchase.Timestamp = time.Now()
		}
	}
	if purchase.ID == "" {
		return nil, errors.New("no products in lineItems")
	}

	if customer != nil {
		purchase.Email = customer.Email
		purchase.Name = customer.Name
	}

	return &purchase, nil
}

func firstOrZero(slice []string) string {
	if len(slice) == 0 {
		return ""
	}
	return slice[0]
}
