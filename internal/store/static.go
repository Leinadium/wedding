package store

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
)

type StaticStore struct {
	lock sync.RWMutex
}

func NewStaticStore() *StaticStore {
	return &StaticStore{}
}

func (s *StaticStore) Products(_ context.Context) ([]Product, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return products, nil
}

func (s *StaticStore) Product(_ context.Context, pid ProductID) (Product, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, p := range products {
		if p.StripeID == pid {
			return p, nil
		}
	}
	return Product{}, errors.New("product not found")
}

func (s *StaticStore) NewPurchase(_ context.Context, purchase Purchase) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	purchases = append(purchases, purchase)
	return nil

}

func (s *StaticStore) Purchases(_ context.Context) ([]Purchase, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return purchases, nil
}

func (s *StaticStore) Sync(_ context.Context, active, inactive []Product) error {
	return nil
}

func (s *StaticStore) NewInvite(ctx context.Context, invite Invite) (InviteID, error) {
	return InviteID("AAABBB"), nil
}

func (s *StaticStore) Invite(_ context.Context, inviteID InviteID) (Invite, error) {
	return Invite{}, nil
}

func (s *StaticStore) Invites(_ context.Context) ([]Invite, error) {
	return nil, nil
}

func (s *StaticStore) NewAttendee(ctx context.Context, inviteID InviteID, attendee Attendee) error {
	return nil
}

func (s *StaticStore) Attendee(_ context.Context, attendeeID uuid.UUID) (Attendee, error) {
	return Attendee{}, nil
}

func (s *StaticStore) Attendees(ctx context.Context) ([]Attendee, error) {
	return nil, nil
}

func (s *StaticStore) DeleteAttendee(_ context.Context, attendeeID uuid.UUID) error {
	return nil
}

func (s *StaticStore) UpsertAttendee(ctx context.Context, attendee Attendee) error {
	return nil
}

func (s *StaticStore) UpsertNoteInvite(ctx context.Context, inviteID InviteID, note string) error {
	return nil
}

func (s *StaticStore) DeleteInvite(_ context.Context, inviteID InviteID) error {
	return nil
}

var (
	products = []Product{
		{
			StripeID:  "prod_UZ7qFL3mYeEotv",
			Name:      "First product",
			ImageURL:  "1.png",
			PriceBRL:  1234,
			Purchased: false,
		}, {
			StripeID:  "2",
			Name:      "Second",
			ImageURL:  "1.png",
			PriceBRL:  3224,
			Purchased: false,
		}, {
			StripeID:  "3",
			Name:      "Third product very long name such very long name",
			ImageURL:  "1.png",
			PriceBRL:  12323,
			Purchased: false,
		}, {
			StripeID:  "4",
			Name:      "Product that is purchased",
			ImageURL:  "1.png",
			PriceBRL:  4433,
			Purchased: true,
		}, {
			StripeID:  "5",
			Name:      "Fifth product",
			ImageURL:  "1.png",
			PriceBRL:  999,
			Purchased: false,
		}, {
			StripeID:  "6",
			Name:      "S",
			ImageURL:  "1.png",
			PriceBRL:  1,
			Purchased: true,
		}, {
			StripeID:  "7",
			Name:      "Simple seventh product",
			ImageURL:  "1.png",
			PriceBRL:  12323,
			Purchased: false,
		}, {
			StripeID:  "8",
			Name:      "This product is already purchased and it costs a lot",
			ImageURL:  "1.png",
			PriceBRL:  443322,
			Purchased: true,
		}, {
			StripeID:  "9",
			Name:      "Ninth product wow",
			ImageURL:  "1.png",
			PriceBRL:  78888,
			Purchased: true,
		}, {
			StripeID:  "10",
			Name:      "Tenth product!!",
			ImageURL:  "1.png",
			PriceBRL:  3224,
			Purchased: false,
		}, {
			StripeID:  "11",
			Name:      "Does it support unicode? 💍",
			ImageURL:  "1.png",
			PriceBRL:  112299,
			Purchased: false,
		}, {
			StripeID:  "12",
			Name:      "I hope it does! 🎄",
			ImageURL:  "1.png",
			PriceBRL:  8787,
			Purchased: true,
		},
	}

	purchases = []Purchase{
		{
			ID:          "1",
			Email:       "test@test.com",
			ProductID:   "1",
			ProductName: "First product",
			Price:       1234,
		},
	}
)
