package notification

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
)

const template = "Wedding update: %s"

type Params struct {
	Token  string
	ChatID string
}

type TelegramNotificator struct {
	bot  *bot.Bot
	chat string
}

func NewTelegramNotificator(params Params) (*TelegramNotificator, error) {
	bot, err := bot.New(params.Token)
	if err != nil {
		return nil, err
	}
	return &TelegramNotificator{
		bot:  bot,
		chat: params.ChatID,
	}, nil
}

func (n *TelegramNotificator) Notify(ctx context.Context, msg string) error {
	_, err := n.bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: n.chat,
		Text:   fmt.Sprintf(template, msg),
	})
	return err
}
