package models

import "time"

type Guest struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductID string
type Product struct {
	StripeID  ProductID `json:"stripe_id" gorm:"primary_key"`
	Name      string    `json:"name"`
	ImageURL  string    `json:"image_url"`
	PriceBRL  int64     `json:"price_brl"`
	Purchased bool      `json:"purchased"`
}

type Purchase struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
}

// do not store in table
type Payment struct {
	URL string
}
