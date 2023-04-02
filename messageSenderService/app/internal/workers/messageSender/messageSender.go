package messageSender

import (
	"context"
	"fmt"
	"github.com/Markuysa/astroMSA/astroService/app/pkg/constanses"
	"github.com/Markuysa/astroMSA/astroService/app/pkg/workers/astroWorker"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/config"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/helpers/htmlHelper"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/model"
	modelExternal "github.com/Markuysa/astroMSA/messageSenderService/app/pkg/model"
	"gopkg.in/gomail.v2"
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
	// FIX HERE astroworker
	wg := sync.WaitGroup{}
	bodyPath := "app/ui/Prediction.html"
	var worker astroWorker.AstroWorker
	for _, receiver := range receivers {
		go func(receiver modelExternal.Receiver) {
			wg.Add(1)
			defer wg.Done()
			prediction, _ := worker.FetchPrediction(ctx, constanses.ARIES, "today")
			body, err := htmlHelper.GetHTMLDailyPrediction(bodyPath, *prediction)

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
