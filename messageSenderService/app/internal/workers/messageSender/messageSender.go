package messageSender

import (
	"github.com/hashicorp/go.net/context"
	"messageSenderService/app/internal/config"
)

type MessageSender interface {
	Send(ctx context.Context, receiver string, body string) error
}

type MsgSenderService struct {
	config config.Config
}
