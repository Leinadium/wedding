package store

import (
	"context"

	"github.com/google/uuid"
	"leinadium.dev/wedding/internal/models"
)

type Service interface {
	Products(ctx context.Context) ([]models.Product, error)
	Product(ctx context.Context, pid models.ProductID) (models.Product, error)

	NewPurchase(ctx context.Context, purchase models.Purchase) error
	Purchases(ctx context.Context) ([]models.Purchase, error)

	Sync(ctx context.Context, active, inactive []models.Product) error

	NewInvite(ctx context.Context, invite models.Invite) (models.InviteID, error)
	Invite(ctx context.Context, inviteID models.InviteID) (models.Invite, error)
	UpsertNoteInvite(ctx context.Context, inviteID models.InviteID, note string) error

	NewAttendee(ctx context.Context, inviteID models.InviteID, attendee models.Attendee) error
	Attendees(ctx context.Context) ([]models.Attendee, error)
	Attendee(ctx context.Context, attendeeID uuid.UUID) (models.Attendee, error)
	UpsertAttendee(ctx context.Context, attendee models.Attendee) error
}
