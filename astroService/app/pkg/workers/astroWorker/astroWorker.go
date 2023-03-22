package astroWorker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Markuysa/astroMSA/astroService/app/internal/config"
	"github.com/Markuysa/astroMSA/astroService/app/pkg/model"
	"io/ioutil"
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

	// TODO fix the problem with fastHTTP
	//url := "https://sameer-kumar-aztro-v1.p.rapidapi.com/?sign=%s&day=%s"
	//
	//paramURL := fmt.Sprintf(url, sign, day)
	//
	//req := fasthttp.AcquireRequest()
	//defer fasthttp.ReleaseRequest(req)
	//
	//resp := fasthttp.AcquireResponse()
	//defer fasthttp.ReleaseResponse(resp)
	//
	//req.SetRequestURI(paramURL)
	//
	//req.Header.Add("X-RapidAPI-Key", "54208fc179msh16e7cc7ea2939dcp1eb840jsnde68a4a6c957")
	//req.Header.Add("X-RapidAPI-Host", "sameer-kumar-aztro-v1.p.rapidapi.com")
	//
	//err := fasthttp.Do(req, resp)
	//if err != nil {
	//	return nil, FetchingPredictionError
	//}
	//body := resp.Body()
	//if err != nil {
	//	return nil, err
	//}
	//var prediction model.Prediction
	//fmt.Println(string(body))
	//err = json.Unmarshal(body, &prediction)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &prediction, nil
	url := "https://sameer-kumar-aztro-v1.p.rapidapi.com/?sign=%s&day=%s"

	url = fmt.Sprintf(url, sign, day)

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("X-RapidAPI-Key", w.config.ApiKey)
	req.Header.Add("X-RapidAPI-Host", w.config.ApiHost)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var prediction model.Prediction
	if err := json.Unmarshal(body, &prediction); err != nil {
		return nil, err
	}
	return &prediction, nil
}
