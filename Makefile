build:
	go build -o dist/main cmd/main.go

dev:
	USE_STATIC_STORE=true STRIPE_SECRET="a" STRIPE_KEY="b" \
		go run cmd/main.go