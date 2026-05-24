package models

import (
	"time"

	"github.com/google/uuid"
)

type Confirmation struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	IsChild   bool      `json:"isChild"`
	CreatedAt time.Time
}

type Rejection struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	OtherGuests string    `json:"otherGuests"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ProductID string
type Product struct {
	StripeID  ProductID `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	ImageURL  string    `json:"imageUrl"`
	PriceBRL  int64     `json:"priceBrl"`
	Purchased bool      `json:"purchased"`
}

type Purchase struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	ProductID   string `json:"productId"`
	ProductName string `json:"productName"`
	Price       int64  `json:"price"`
}

// do not store in table
type Payment struct {
	URL string `json:"url"`
}
