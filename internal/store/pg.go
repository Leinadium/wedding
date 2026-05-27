package store

import (
	"context"
	"fmt"
	"math/rand"
	"slices"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"leinadium.dev/wedding/internal/models"
)

type Params struct {
	DSN         string
	AutoMigrate bool
}

type PGStore struct {
	db *gorm.DB
}

func NewPGStore(p Params) Service {
	dsn := p.DSN
	if dsn == "" {
		panic("dsn is required")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	if p.AutoMigrate {
		_ = db.AutoMigrate(&models.Invite{})
		_ = db.AutoMigrate(&models.Attendee{})
		_ = db.AutoMigrate(&models.Product{})
		_ = db.AutoMigrate(&models.Purchase{})
	}

	return &PGStore{
		db: db,
	}
}

func (p *PGStore) Products(ctx context.Context) ([]models.Product, error) {
	return gorm.G[models.Product](p.db).Find(ctx)
}

func (p *PGStore) Product(ctx context.Context, pid models.ProductID) (models.Product, error) {
	return gorm.G[models.Product](p.db).Where("stripe_id = ?", pid).First(ctx)
}

func (p *PGStore) NewPurchase(ctx context.Context, purchase models.Purchase) error {
	return gorm.G[models.Purchase](p.db).Create(ctx, &purchase)
}

func (p *PGStore) Purchases(ctx context.Context) ([]models.Purchase, error) {
	return gorm.G[models.Purchase](p.db).Find(ctx)
}

func (p *PGStore) Sync(ctx context.Context, active, inactive []models.Product) error {
	final := slices.Concat(active, inactive)

	return p.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&final).Error
}

func (p *PGStore) NewInvite(ctx context.Context, invite models.Invite) (models.InviteID, error) {
	inviteID := generateInviteID()
	var success bool

	for range 3 {
		invite.ID = inviteID
		if err := gorm.G[models.Invite](p.db).Create(ctx, &invite); err == nil {
			success = true
			break
		}
		inviteID = generateInviteID()
	}
	if !success {
		return "", fmt.Errorf("could not generate invite ID")
	}

	return inviteID, nil
}

func (p *PGStore) Invite(ctx context.Context, inviteID models.InviteID) (models.Invite, error) {
	invite, err := gorm.G[models.Invite](p.db).Preload("Attendee", nil).Where("id = ?", inviteID).First(ctx)
	if err != nil {
		return models.Invite{}, err
	}
	return invite, nil
}

func (p *PGStore) UpsertNoteInvite(ctx context.Context, inviteID models.InviteID, note string) error {
	_, err := gorm.G[models.Invite](p.db).Where("id = ?", inviteID).Update(ctx, "note", note)
	return err
}

func (p *PGStore) DeleteInvite(ctx context.Context, inviteID models.InviteID) error {
	_, err := gorm.G[models.Invite](p.db).Where("id = ?", inviteID).Delete(ctx)
	return err
}

func (s *PGStore) NewAttendee(ctx context.Context, inviteID models.InviteID, attendee models.Attendee) error {
	attendee.InviteID = inviteID
	return gorm.G[models.Attendee](s.db).Create(ctx, &attendee)
}

func (p *PGStore) Attendee(ctx context.Context, attendeeID uuid.UUID) (models.Attendee, error) {
	return gorm.G[models.Attendee](p.db).Where("id = ?", attendeeID).First(ctx)
}

func (p *PGStore) Attendees(ctx context.Context) ([]models.Attendee, error) {
	return gorm.G[models.Attendee](p.db).Find(ctx)
}

func (p *PGStore) UpsertAttendee(ctx context.Context, attendee models.Attendee) error {
	return gorm.G[models.Attendee](p.db).Create(ctx, &attendee)
}

func (p *PGStore) DeleteAttendee(ctx context.Context, attendeeID uuid.UUID) error {
	_, err := gorm.G[models.Attendee](p.db).Where("id = ?", attendeeID).Delete(ctx)
	return err
}

func generateInviteID() models.InviteID {
	id := make([]byte, 6)
	for i := range id {
		id[i] = idCharset[rand.Intn(len(idCharset))]
	}
	return models.InviteID(id)
}

const idCharset = "23456789ABCDEFGHJKMNPQRSTUVWXYZ"
