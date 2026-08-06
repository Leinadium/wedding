package store

import (
	"context"

	"github.com/google/uuid"
)

type Store interface {
	Products(ctx context.Context) ([]Product, error)
	Product(ctx context.Context, pid ProductID) (Product, error)

	NewPurchase(ctx context.Context, purchase Purchase) error
	Purchases(ctx context.Context) ([]Purchase, error)

	Sync(ctx context.Context, active, inactive []Product) error

	NewInvite(ctx context.Context, invite Invite) (InviteID, error)
	Invite(ctx context.Context, inviteID InviteID) (Invite, error)
	Invites(ctx context.Context) ([]Invite, error)
	UpsertNoteInvite(ctx context.Context, inviteID InviteID, note string) error
	DeleteInvite(ctx context.Context, inviteID InviteID) error

	NewAttendee(ctx context.Context, inviteID InviteID, attendee Attendee) error
	Attendees(ctx context.Context) ([]Attendee, error)
	Attendee(ctx context.Context, attendeeID uuid.UUID) (Attendee, error)
	UpsertAttendee(ctx context.Context, attendee Attendee) error
	DeleteAttendee(ctx context.Context, attendeeID uuid.UUID) error
}
