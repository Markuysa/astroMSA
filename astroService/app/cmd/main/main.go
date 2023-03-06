package main

import (
	config2 "astroService/internal/config"
	"astroService/pkg/constanses"
	astroWorker "astroService/pkg/workers/astroWorker"
	"fmt"
	"github.com/hashicorp/go.net/context"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

// checking the response
func main() {

	ctx := context.Background()
	config, err := config2.New()
	if err != nil {
		log.Fatal(err)
	}
	client := fasthttp.Client{ReadTimeout: 5 * time.Second}
	astrologyWorker := astroWorker.Init(config, &client)
	prediction, err := astrologyWorker.FetchPrediction(ctx, constanses.ARIES, "today")
	if err != nil {
		return
	}
	fmt.Println(prediction)
}
