# wedding

My own wedding website. To be hosted on wedding.leinadium.dev

## API

The API is a Go project, using the Gin framework. Has integrations with:

- Stripe for payment processing and product listing
- PostgreSQL for data storage
- Telegram for notifications

There is a CLI tool for administrative tasks that interact with the API. It can:

- Manage invites
- Manage attendees
- List products
- List purchases

All the API code was written manually (except the boilerplate that is the internal/client/request.go file).

Refer to the `internal/v1` package for features, and `internal/server` for the HTTP server implementation.
Database schema can be found in the `internal/models` package.

### Executing

Add necessary environment variables to the `.env` file, following the `example.env` file.

Use Docker Compose to run the application, with:

```bash
docker compose up --build
```

It builds the API and spawns a PostgreSQL instance.

Use the CLI to manage invites, attendees, products, and purchases. To run:

```bash
# first, install it
go install ./cmd/wedding

wedding -h
```
