package astroWorker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Markuysa/astroMSA/astroService/app/internal/config"
	"github.com/Markuysa/astroMSA/astroService/app/pkg/constanses"
	"github.com/Markuysa/astroMSA/astroService/app/pkg/model"
	"github.com/valyala/fasthttp"
	"net/http"
)

var (
	FetchingPredictionError = errors.New("error with fetching the prediction")
)

type AstroFetcher interface {
	FetchPrediction(ctx context.Context, sign string, day string) (*model.Prediction, error)
}
type AstroWorker struct {
	config     *config.Config
	httpClient *http.Client
}

func Init(config *config.Config, client *http.Client) *AstroWorker {
	return &AstroWorker{
		config:     config,
		httpClient: client,
	}
}
func (w *AstroWorker) FetchPrediction(ctx context.Context, sign string, day string) (*model.Prediction, error) {
	url := "https://sameer-kumar-aztro-v1.p.rapidapi.com/?sign=%s&day=%s"

	paramURL := fmt.Sprintf(url, sign, day)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(paramURL)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.Add("X-RapidAPI-Key", "54208fc179msh16e7cc7ea2939dcp1eb840jsnde68a4a6c957")
	req.Header.Add("X-RapidAPI-Host", "sameer-kumar-aztro-v1.p.rapidapi.com")
	req.Header.SetMethod(fasthttp.MethodPost)

	err := fasthttp.Do(req, resp)
	if err != nil {
		return nil, FetchingPredictionError
	}
	body := resp.Body()
	var prediction model.Prediction
	err = json.Unmarshal(body, &prediction)
	if err != nil {
		return nil, err
	}
	if len(prediction.DateRange) == 0 {
		return constanses.ReturnConst(sign)
	}
	return &prediction, nil
}
