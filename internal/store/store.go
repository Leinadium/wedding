package store

import (
	"context"

	"leinadium.dev/wedding/internal/models"
)

type Service interface {
	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProduct(ctx context.Context, pid models.ProductID) (models.Product, error)
	CreateGuest(ctx context.Context, guest models.Guest) error
	CreatePurchase(ctx context.Context, purchase models.Purchase) error
	GetPurchases(ctx context.Context) ([]models.Purchase, error)

	Sync(ctx context.Context, active, inactive []models.Product) error
}
