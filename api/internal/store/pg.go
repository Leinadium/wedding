package store

import (
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := p.DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if p.AutoMigrate {
		_ = db.AutoMigrate(&models.Guest{})
		_ = db.AutoMigrate(&models.Product{})
		_ = db.AutoMigrate(&models.Purchase{})
	}

	return &PGStore{
		db: db,
	}
}

func (p *PGStore) GetProducts(ctx context.Context) ([]models.Product, error) {
	return gorm.G[models.Product](p.db).Find(ctx)
}

func (p *PGStore) GetProduct(ctx context.Context, pid models.ProductID) (models.Product, error) {
	return gorm.G[models.Product](p.db).Where("id = ?", pid).First(ctx)
}

func (p *PGStore) CreateGuest(ctx context.Context, guest models.Guest) error {
	return gorm.G[models.Guest](p.db).Create(ctx, &guest)
}

func (p *PGStore) CreatePurchase(ctx context.Context, purchase models.Purchase) error {
	return gorm.G[models.Purchase](p.db).Create(ctx, &purchase)
}

func (p *PGStore) GetPurchases(ctx context.Context) ([]models.Purchase, error) {
	return gorm.G[models.Purchase](p.db).Find(ctx)
}
