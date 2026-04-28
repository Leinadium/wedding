package internal

import "os"

type Params struct {
	ServerURL           string
	StripeKey           string
	StripeSecret        string
	AuthSecret          string
	DatabaseDSN         string
	DatabaseAutomigrate bool
	Port                int
}

func NewParams() Params {
	return Params{
		ServerURL:           os.Getenv("SERVER_URL"),
		StripeKey:           os.Getenv("STRIPE_KEY"),
		StripeSecret:        os.Getenv("STRIPE_SECRET"),
		AuthSecret:          os.Getenv("AUTH_SECRET"),
		DatabaseDSN:         os.Getenv("DATABASE_DSN"),
		DatabaseAutomigrate: os.Getenv("DATABASE_AUTOMIGRATE") == "true",
		Port:                8080,
	}
}
