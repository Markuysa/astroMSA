package astroWorker

import (
	"astroService/pkg/model"
	"errors"
	"github.com/hashicorp/go.net/context"
	"github.com/irfansofyana/go-aztro-api-wrapper/aztro"
)

var (
	FetchingPredictionError = errors.New("error with fetching the prediction")
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

	requestParams := aztro.NewAztroRequestParam(sign, aztro.WithDay(day))

	response, err := w.APIclient.GetHoroscope(requestParams)
	if err != nil {
		return nil, FetchingPredictionError
	}
	return &model.Prediction{
		Color:         response.Color,
		Compatibility: response.Compatibility,
		Description:   response.Description,
		LuckyNumber:   response.LuckyNumber,
		LuckyTime:     response.LuckyTime,
		Mood:          response.Mood,
	}, nil

}
