package v1

import "leinadium.dev/wedding/internal/store"

type Payment struct {
	URL string `json:"url"`
}

type PurchasableProduct struct {
	store.Product
	Purchased bool `json:"purchased"`
}
