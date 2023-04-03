package messageSender

import (
	"context"
	"fmt"
	astroModels "github.com/Markuysa/astroMSA/astroService/app/pkg/model"
	"github.com/Markuysa/astroMSA/messageSenderService/app/gapi/client"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/config"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/helpers/htmlHelper"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/model"
	modelExternal "github.com/Markuysa/astroMSA/messageSenderService/app/pkg/model"
	"github.com/Markuysa/astroMSA/messageSenderService/app/protobuf/pb"
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

func (s *MsgSenderWorker) SendDailyPredictions(ctx context.Context, req *pb.DailyPredictionsRequest) error {

	wg := sync.WaitGroup{}
	bodyPath := "app/ui/Prediction.html"
	receiversChan := make(chan astroModels.HandledPrediction, 10)
	// TODO idk how to fix this stuff )()(()O
	for _, receiver := range req.Receivers {
		go func(receiver any) {
			client.FetchPrediction(ctx,receiver.)

		}(receiver)
		go func(receiver modelExternal.Receiver) {
			wg.Add(1)
			defer wg.Done()

			body, err := htmlHelper.GetHTMLDailyPrediction(bodyPath)

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
