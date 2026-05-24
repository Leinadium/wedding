package store

import (
	"context"
	"fmt"
	"slices"

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

func (p *PGStore) NewConfirmations(ctx context.Context, confirmations []models.Confirmation) error {
	return gorm.G[models.Confirmation](p.db).CreateInBatches(ctx, &confirmations, 5)
}

func (p *PGStore) NewRejection(ctx context.Context, rejection models.Rejection) error {
	return gorm.G[models.Rejection](p.db).Create(ctx, &rejection)
}
