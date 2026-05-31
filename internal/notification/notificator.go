package notification

import "context"

type Notificator interface {
	Notify(ctx context.Context, msg string) error
}
