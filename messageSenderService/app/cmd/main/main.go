package main

import (
	"github.com/hashicorp/go.net/context"
	"log"
	"messageSenderService/app/internal/config"
	"messageSenderService/app/internal/workers/messageSender"
	"messageSenderService/app/pkg/model"
)

func main() {

	ctx := context.Background()
	messageConfig, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	messageWorker := messageSender.Init(messageConfig)

	err = messageWorker.SendDailyPredictions(ctx, []model.Receiver{
		model.Receiver{Email: "islam.kus1111@gmail.com", Zodiac: "Leo"},
		model.Receiver{Email: "islam.kus1111@gmail.com", Zodiac: "Leo"},
		model.Receiver{Email: "islam.kus1111@gmail.com", Zodiac: "Leo"},
		model.Receiver{Email: "islam.kus1111@gmail.com", Zodiac: "Leo"},
		model.Receiver{Email: "markuysa.study@mail.ru", Zodiac: "Aries"}})
	if err != nil {
		log.Print(err)
	}

	//cron := cronWorker.Init()
	//cron.StartTicker()
}
