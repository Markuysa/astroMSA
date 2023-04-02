package astroWorker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Markuysa/astroMSA/astroService/app/internal/config"
	"github.com/Markuysa/astroMSA/astroService/app/pkg/model"
	"io/ioutil"
	"log"
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
	log.Println(sign, day)
	url = fmt.Sprintf(url, sign, day)

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("X-RapidAPI-Key", w.config.ApiKey)
	req.Header.Add("X-RapidAPI-Host", w.config.ApiHost)

	log.Println(req)
	res, _ := http.DefaultClient.Do(req)
	log.Print(res)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var prediction model.Prediction
	if err := json.Unmarshal(body, &prediction); err != nil {
		return nil, err
	}
	return &prediction, nil
}
