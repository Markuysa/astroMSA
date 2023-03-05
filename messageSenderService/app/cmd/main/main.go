package main

import "messageSenderService/app/internal/workers/cronWorker"

func main() {

	//ctx := context.Background()
	//messageConfig, err := config.Init()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//messageWorker := messageSender.Init(messageConfig)
	//
	//err = messageWorker.SendDailyPredictions(ctx, []model.Receiver{
	//	model.Receiver{Email: "islam.kus1111@gmail.com", Zodiac: "Leo"},
	//	model.Receiver{Email: "islam.kus1111@gmail.com", Zodiac: "Leo"},
	//	model.Receiver{Email: "islam.kus1111@gmail.com", Zodiac: "Leo"},
	//	model.Receiver{Email: "islam.kus1111@gmail.com", Zodiac: "Leo"},
	//	model.Receiver{Email: "markuysa.study@mail.ru", Zodiac: "Aries"}})
	//if err != nil {
	//	log.Print(err)
	//}

	cron := cronWorker.Init()
	cron.StartTicker()
}
