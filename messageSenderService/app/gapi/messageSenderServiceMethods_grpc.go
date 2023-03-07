package gapi

import (
	"context"
	"messageSenderService/app/protobuf/pb"
)

func (s *Server) SendDailyPredictions(context.Context, *pb.DailyPredictionsRequest) (*pb.DailyPredictionsResponse, error) {
	return nil, nil
}
