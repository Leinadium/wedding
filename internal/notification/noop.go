package notification

import "context"

type NoopNotificator struct{}

func NewNoopNotificator() *NoopNotificator {
	return &NoopNotificator{}
}

func (n *NoopNotificator) Notify(ctx context.Context, msg string) error {
	return nil
}
