package internal

import "github.com/caarlos0/env/v11"

type Params struct {
	ServerURL  string `env:"SERVER_URL" envDefault:"http://localhost:8080"`
	AuthSecret string `env:"AUTH_SECRET" envDefault:"changeMe!"`
	Port       int    `env:"PORT" envDefault:"8080"`

	StripeKey    string `env:"STRIPE_KEY,notEmpty"`
	StripeSecret string `env:"STRIPE_SECRET,notEmpty"`

	DatabaseDSN         string `env:"DATABASE_DSN"`
	DatabaseAutomigrate bool   `env:"DATABASE_AUTOMIGRATE" envDefault:"true"`

	TelegramToken  string `env:"TELEGRAM_TOKEN"`
	TelegramChatID string `env:"TELEGRAM_CHAT_ID"`

	UseStaticStore bool   `env:"USE_STATIC_STORE" envDefault:"false"`
	UseSync        bool   `env:"USE_SYNC" envDefault:"false"`
	StaticDir      string `env:"STATIC_DIR"`
}

func NewParams() (p Params, err error) {
	err = env.Parse(&p)
	return
}
