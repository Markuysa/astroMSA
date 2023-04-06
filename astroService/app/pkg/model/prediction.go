package model

import "github.com/Markuysa/astroMSA/astroService/app/protobuf/pb"

// Prediction structure of astro prediction
type Prediction struct {
	DateRange     string `json:"date_range"`
	CurrentDate   string `json:"current_date"`
	Description   string `json:"description"`
	Compatibility string `json:"compatibility"`
	Mood          string `json:"mood"`
	Color         string `json:"color"`
	LuckyNumber   string `json:"lucky_number"`
	LuckyTime     string `json:"lucky_time"`
}

type HandledPrediction struct {
	Prediction  *pb.PredictionResponse
	Destination string
}
