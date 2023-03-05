package astroWorker

import (
	"astroService/pkg/model"
	"github.com/hashicorp/go.net/context"
	"github.com/irfansofyana/go-aztro-api-wrapper/aztro"
)

type AstroFetcher interface {
	FetchPrediction(ctx context.Context, sign aztro.Sign, day aztro.Day) (*model.Prediction, error)
}
type AstroWorker struct {
	APIclient *aztro.AztroClient
}

func Init(client *aztro.AztroClient) *AstroWorker {
	return &AstroWorker{
		APIclient: client,
	}
}
func (w *AstroWorker) FetchPrediction(ctx context.Context, sign aztro.Sign, day aztro.Day) (*model.Prediction, error) {

	w.APIclient.GetHoroscope()
}
