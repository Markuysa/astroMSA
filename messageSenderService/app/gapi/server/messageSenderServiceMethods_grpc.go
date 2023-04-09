package server

import (
	"context"
	pb "github.com/Markuysa/astroMSA/apigw/app/protobuf/gen"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/helpers/protoHelper"
)

// SendDailyPredictions method is needed to
// handle /api/v1/sendPredictions request and send
// daily predictions to users
func (s *Server) SendDailyPredictions(ctx context.Context, req *pb.DailyPredictionsRequest) (*pb.DailyPredictionsResponse, error) {

	receivers := protoHelper.ReceiversFromPb(req)
	err := s.MessageWorker.SendDailyPredictions(ctx, receivers)
	if err != nil {
		return &pb.DailyPredictionsResponse{ResponseStatus: false}, err
	}
	return &pb.DailyPredictionsResponse{ResponseStatus: true}, nil
}
