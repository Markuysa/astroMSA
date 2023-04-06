package htmlHelper

import (
	"bytes"
	"errors"
	"github.com/Markuysa/astroMSA/astroService/app/pkg/model"
	"html/template"
)

var (
	ParseHTMLError   = errors.New("parse html error")
	ExecuteHTMLError = errors.New("execute html error (watch the struct)")
)

// GetHTMLDailyPrediction method parses the HTML template and fills
// it with necessary data from prediction object
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
