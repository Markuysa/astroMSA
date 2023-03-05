package messageSender

import (
	astroModel "astroService/pkg/model"
	"fmt"
	"github.com/hashicorp/go.net/context"
	"gopkg.in/gomail.v2"
	"messageSenderService/app/internal/config"
	"messageSenderService/app/internal/helpers/htmlHelper"
	"messageSenderService/app/internal/model"
	modelExternal "messageSenderService/app/pkg/model"
	"sync"
)

type MessageSender interface {
	SendHTMl(ctx context.Context, message model.Message) error
	SendDailyPredictions(ctx context.Context, receivers []string) error
}

type MsgSenderWorker struct {
	config *config.Config
}

func Init(config *config.Config) *MsgSenderWorker {
	return &MsgSenderWorker{
		config: config,
	}
}

func (s *MsgSenderWorker) SendHTML(ctx context.Context, message model.Message) error {

	m := gomail.NewMessage()
	m.SetHeader("From", s.config.MessageSenderMail)
	m.SetHeader("To", message.Receiver)
	m.SetHeader("Subject", message.Subject)
	m.SetBody("text/html", message.Body)
	d := gomail.NewDialer("smtp.gmail.com", 587, s.config.MessageSenderMail, s.config.MessageSenderPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
	return nil
}

func (s *MsgSenderWorker) SendDailyPredictions(ctx context.Context, receivers []modelExternal.Receiver) error {

	wg := sync.WaitGroup{}
	bodyPath := "app/ui/confirm-email.html"
	for _, receiver := range receivers {
		go func(receiver modelExternal.Receiver) {
			wg.Add(1)
			defer wg.Done()
			body, err := htmlHelper.GetHTMLDailyPrediction(bodyPath, astroModel.Prediction{})

			err = s.SendHTML(ctx, model.Message{
				Receiver: receiver.Email,
				Body:     body.String(),
				Subject:  "Daily prediction",
			})
			if err != nil {
				return
			}
		}(receiver)
	}
	wg.Wait()
	return nil
}
