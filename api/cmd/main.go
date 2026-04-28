package main

import (
	"fmt"

	"leinadium.dev/wedding/internal"
	"leinadium.dev/wedding/internal/payment"
	"leinadium.dev/wedding/internal/server"
	"leinadium.dev/wedding/internal/store"
	v1 "leinadium.dev/wedding/internal/v1"
)

func main() {
	// creating params
	p := internal.NewParams()

	// creating services
	storeService := store.NewPGStore(store.Params{
		DSN:         p.DatabaseDSN,
		AutoMigrate: p.DatabaseAutomigrate,
	})
	paymentService := payment.New(payment.Params{
		Domain:        p.ServerURL,
		Key:           p.StripeKey,
		WebhookSecret: p.StripeSecret,
	})
	v1Service := v1.New(storeService, paymentService, v1.Params{})

	// creating server
	sv := server.New(v1Service, server.Params{
		AuthSecret: p.StripeSecret,
	})

	err := sv.Run(p.Port)
	fmt.Println(err)
}
