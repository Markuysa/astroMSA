package helpers

import (
	"github.com/Markuysa/astroMSA/astroService/app/pkg/model"
	"github.com/Markuysa/astroMSA/astroService/app/protobuf/pb"
)

// PredictionToPb is the converter - from default
// project Prediction object to protobuf object
func PredictionToPb(prediction *model.Prediction) *pb.PredictionResponse {

	return &pb.PredictionResponse{
		DateRange:     prediction.DateRange,
		CurrentDate:   prediction.CurrentDate,
		Description:   prediction.Description,
		Compatibility: prediction.Compatibility,
		Mood:          prediction.Mood,
		Color:         prediction.Color,
		LuckyNumber:   prediction.LuckyNumber,
		LuckyTime:     prediction.LuckyTime,
	}

}
