package messageSender

import (
	"context"
	"fmt"
	pb "github.com/Markuysa/astroMSA/apiGateway/app/protobuf/gen"
	astroModels "github.com/Markuysa/astroMSA/astroService/app/pkg/model"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/config"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/helpers/htmlHelper"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/helpers/protoHelper"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/model"
	externalModels "github.com/Markuysa/astroMSA/messageSenderService/app/pkg/model"
	"gopkg.in/gomail.v2"
	"log"
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

// SendHTML methods sends a message using gmail
// with some HTML template
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

// SendDailyPredictions sends predictions to users in goroutines
// Maybe should receive slice with emails of receivers than struct ?>?
func (s *MsgSenderWorker) SendDailyPredictions(ctx context.Context, req []externalModels.Receiver) error {
	wg := sync.WaitGroup{}
	bodyPath := "app/ui/Prediction.html"
	receiversChan := make(chan astroModels.HandledPrediction, 10)
	log.Println(req)
	go func() {
		for _, receiver := range req {
			//pbPrediction, err := client.FetchPrediction(ctx, "today", receiver.Zodiac)
			pbPrediction := pb.PredictionResponse{
				Mood:          "smth",
				LuckyTime:     "som",
				DateRange:     "asd",
				Description:   "u will be the happiest man in the world)",
				CurrentDate:   "today))",
				Compatibility: "aries",
				Color:         "black",
				LuckyNumber:   "1",
			}
			prediction := protoHelper.PredictionFromPb(&pbPrediction)
			//if err != nil {
			//	return
			//}
			receiversChan <- astroModels.HandledPrediction{Prediction: prediction, Destination: receiver.Email}
		}
		close(receiversChan)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for receiver := range receiversChan {
			body, err := htmlHelper.GetHTMLDailyPrediction(bodyPath, *receiver.Prediction)
			err = s.SendHTML(ctx, model.Message{
				Receiver: receiver.Destination,
				Body:     body.String(),
				Subject:  "Daily prediction",
			})
			if err != nil {
				return
			}
		}
	}()
	wg.Wait()
	log.Println("ENDED")
	return nil
}
