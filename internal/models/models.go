package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null/v6"
)

type InviteID string
type Invite struct {
	ID        InviteID   `gorm:"primary_key" json:"id"`
	Phone     string     `json:"phone"`
	Note      string     `json:"note"`
	Attendees []Attendee `json:"attendees" gorm:"foreignKey:InviteID;constraint:OnDelete:CASCADE;"`
}

type AttendeeID = uuid.UUID
type Attendee struct {
	ID        *AttendeeID `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"id"`
	InviteID  InviteID    `json:"inviteId"`
	Name      string      `json:"name"`
	IsChild   bool        `json:"isChild"`
	Confirmed null.Bool   `json:"confirmed"`
	UpdatedAt time.Time   `json:"updatedAt"`
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
