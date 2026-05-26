package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type InviteID string
type Invite struct {
	ID     InviteID `gorm:"primary_key"`
	Phone  string   `json:"phone"`
	Number int      `json:"number"`
	Names  string   `json:"names"`
	Note   string   `json:"note"`
}

type Attendee struct {
	ID        uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid()"`
	InviteID  InviteID     `json:"inviteId"`
	Name      string       `json:"name"`
	IsChild   bool         `json:"isChild"`
	Confirmed sql.NullBool `json:"confirmed"`
	UpdatedAt time.Time    `json:"updatedAt"`
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
