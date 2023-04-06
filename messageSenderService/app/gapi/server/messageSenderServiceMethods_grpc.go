package server

import (
	"context"
	"github.com/Markuysa/astroMSA/messageSenderService/app/protobuf/pb"
)

// SendDailyPredictions method is needed to
// handle /api/v1/sendPredictions request and send
// daily predictions to users
func (s *Server) SendDailyPredictions(ctx context.Context, req *pb.DailyPredictionsRequest) (*pb.DailyPredictionsResponse, error) {

	err := s.MessageWorker.SendDailyPredictions(ctx, req)
	if err != nil {
		return &pb.DailyPredictionsResponse{ResponseStatus: false}, err
	}
	return &pb.DailyPredictionsResponse{ResponseStatus: true}, nil
}
