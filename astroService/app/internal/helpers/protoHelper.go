package helpers

import (
	pb "github.com/Markuysa/astroMSA/apiGateway/app/protobuf/gen"
	"github.com/Markuysa/astroMSA/astroService/app/pkg/model"
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
func PredictionFromPb(prediction *pb.PredictionResponse) *model.Prediction {

	return &model.Prediction{
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
