package protoHelper

import (
	pb "github.com/Markuysa/astroMSA/apigw/app/protobuf/gen"
	"github.com/Markuysa/astroMSA/astroService/app/pkg/model"
	external "github.com/Markuysa/astroMSA/messageSenderService/app/pkg/model"
)

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
func ReceiversFromPb(request *pb.DailyPredictionsRequest) []external.Receiver {
	var receivers []external.Receiver
	for _, receiverpb := range request.Receivers {
		receiver := external.Receiver{Email: receiverpb.Email, Zodiac: receiverpb.Email}
		receivers = append(receivers, receiver)
	}
	return receivers
}
