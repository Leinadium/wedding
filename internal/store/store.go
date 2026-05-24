package store

import (
	"context"

	"leinadium.dev/wedding/internal/models"
)

type Service interface {
	Products(ctx context.Context) ([]models.Product, error)
	Product(ctx context.Context, pid models.ProductID) (models.Product, error)

	NewPurchase(ctx context.Context, purchase models.Purchase) error
	Purchases(ctx context.Context) ([]models.Purchase, error)

	Sync(ctx context.Context, active, inactive []models.Product) error

	NewConfirmations(ctx context.Context, confirmations []models.Confirmation) error
	NewRejection(ctx context.Context, rejection models.Rejection) error
}
