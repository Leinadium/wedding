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
)

type Params struct {
	DSN         string
	AutoMigrate bool
}

type PGStore struct {
	db *gorm.DB
}

func NewPGStore(p Params) Store {
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
		_ = db.AutoMigrate(&Invite{})
		_ = db.AutoMigrate(&Attendee{})
		_ = db.AutoMigrate(&Product{})
		_ = db.AutoMigrate(&Purchase{})
	}

	return &PGStore{
		db: db,
	}
}

func (p *PGStore) Products(ctx context.Context) ([]Product, error) {
	return gorm.G[Product](p.db).Find(ctx)
}

func (p *PGStore) Product(ctx context.Context, pid ProductID) (Product, error) {
	return gorm.G[Product](p.db).Where("stripe_id = ?", pid).First(ctx)
}

func (p *PGStore) NewPurchase(ctx context.Context, purchase Purchase) error {
	return p.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&purchase).Error
}

func (p *PGStore) Purchases(ctx context.Context) ([]Purchase, error) {
	return gorm.G[Purchase](p.db).Find(ctx)
}

func (p *PGStore) Sync(ctx context.Context, active, inactive []Product) error {
	final := slices.Concat(active, inactive)

	return p.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&final).Error
}

func (p *PGStore) NewInvite(ctx context.Context, invite Invite) (InviteID, error) {
	inviteID := generateInviteID()
	var success bool

	for range 3 {
		invite.ID = inviteID
		if err := gorm.G[Invite](p.db).Create(ctx, &invite); err == nil {
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

func (p *PGStore) Invite(ctx context.Context, inviteID InviteID) (Invite, error) {
	invite, err := gorm.G[Invite](p.db).Preload("Attendees", nil).Where("id = ?", inviteID).First(ctx)
	if err != nil {
		return Invite{}, err
	}
	return invite, nil
}

func (p *PGStore) Invites(ctx context.Context) ([]Invite, error) {
	invites, err := gorm.G[Invite](p.db).Preload("Attendees", nil).Find(ctx)
	if err != nil {
		return nil, err
	}
	return invites, nil
}

func (p *PGStore) UpsertNoteInvite(ctx context.Context, inviteID InviteID, note string) error {
	_, err := gorm.G[Invite](p.db).Where("id = ?", inviteID).Update(ctx, "note", note)
	return err
}

func (p *PGStore) DeleteInvite(ctx context.Context, inviteID InviteID) error {
	_, err := gorm.G[Invite](p.db).Where("id = ?", inviteID).Delete(ctx)
	return err
}

func (s *PGStore) NewAttendee(ctx context.Context, inviteID InviteID, attendee Attendee) error {
	attendee.InviteID = inviteID
	return gorm.G[Attendee](s.db).Create(ctx, &attendee)
}

func (p *PGStore) Attendee(ctx context.Context, attendeeID uuid.UUID) (Attendee, error) {
	return gorm.G[Attendee](p.db).Where("id = ?", attendeeID).First(ctx)
}

func (p *PGStore) Attendees(ctx context.Context) ([]Attendee, error) {
	return gorm.G[Attendee](p.db).Find(ctx)
}

func (p *PGStore) UpsertAttendee(ctx context.Context, attendee Attendee) error {
	return p.db.Save(&attendee).Error
}

func (p *PGStore) DeleteAttendee(ctx context.Context, attendeeID uuid.UUID) error {
	_, err := gorm.G[Attendee](p.db).Where("id = ?", attendeeID).Delete(ctx)
	return err
}

func generateInviteID() InviteID {
	id := make([]byte, 6)
	for i := range id {
		id[i] = idCharset[rand.Intn(len(idCharset))]
	}
	return InviteID(id)
}

const idCharset = "23456789ABCDEFGHJKMNPQRSTUVWXYZ"
