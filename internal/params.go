package internal

import "github.com/caarlos0/env/v11"

type Params struct {
	ServerURL           string `env:"SERVER_URL" envDefault:"http://localhost:8080"`
	StripeKey           string `env:"STRIPE_KEY,notEmpty"`
	StripeSecret        string `env:"STRIPE_SECRET,notEmpty"`
	AuthSecret          string `env:"AUTH_SECRET" envDefault:"changeMe!"`
	DatabaseDSN         string `env:"DATABASE_DSN"`
	DatabaseAutomigrate bool   `env:"DATABASE_AUTOMIGRATE" envDefault:"true"`
	Port                int    `env:"PORT" envDefault:"8080"`

	UseStaticStore bool   `env:"USE_STATIC_STORE" envDefault:"false"`
	StaticDir      string `env:"STATIC_DIR"`
}

func NewParams() (p Params, err error) {
	err = env.Parse(&p)
	return
}
