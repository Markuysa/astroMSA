package htmlHelper

import (
	"astroService/app/pkg/model"
	"bytes"
	"errors"
	"html/template"
)

var (
	ParseHTMLError   = errors.New("parse html error")
	ExecuteHTMLError = errors.New("execute html error (watch the struct)")
)

func GetHTMLDailyPrediction(filePath string, prediction model.Prediction) (*bytes.Buffer, error) {

	var body bytes.Buffer
	t, err := template.ParseFiles(filePath)
	if err != nil {
		return nil, ParseHTMLError
	}
	err = t.Execute(&body, prediction)
	if err != nil {
		return nil, ExecuteHTMLError
	}
	return &body, nil
}
