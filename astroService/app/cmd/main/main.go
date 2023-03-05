package main

import (
	astroWorker "astroService/pkg/workers/astroWorker"
	"github.com/hashicorp/go.net/context"
	"github.com/irfansofyana/go-aztro-api-wrapper/aztro"
	"log"
)

// checking the response
func main() {

	ctx := context.Background()
	aztroClient, err := aztro.NewAztroClient()

	if err != nil {
		log.Fatal(err)
	}
	astrologyWorker := astroWorker.Init(aztroClient)

}
