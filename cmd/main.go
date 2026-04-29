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
	p, err := internal.NewParams()
	if err != nil {
		panic(err)
	}

	// creating services
	// store
	var storeService store.Service
	if p.UseStaticStore {
		storeService = store.NewStaticStore()
	} else {
		storeService = store.NewPGStore(store.Params{
			DSN:         p.DatabaseDSN,
			AutoMigrate: p.DatabaseAutomigrate,
		})
	}
	// payment
	paymentService := payment.New(payment.Params{
		Domain:        p.ServerURL,
		Key:           p.StripeKey,
		WebhookSecret: p.StripeSecret,
	})
	// api
	v1Service := v1.New(storeService, paymentService, v1.Params{})

	// creating server and starting
	sv := server.New(v1Service, server.Params{
		AuthSecret: p.StripeSecret,
	})
	err = sv.Run(p.Port)
	fmt.Println(err)
}
