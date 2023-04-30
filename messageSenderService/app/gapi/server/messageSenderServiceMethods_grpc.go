package server

import (
	"context"
	pb "github.com/Markuysa/astroMSA/apiGateway/app/protobuf/gen"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/helpers/protoHelper"
	"go.uber.org/zap"
)

// SendDailyPredictions method is needed to
// handle /api/v1/sendPredictions request and send
// daily predictions to users
func (s *Server) SendDailyPredictions(ctx context.Context, req *pb.DailyPredictionsRequest) (*pb.DailyPredictionsResponse, error) {
	s.Logger.Info("send daily prediction request received", zap.String("day", req.Day))
	receivers := protoHelper.ReceiversFromPb(req)
	go func() {
		err := s.MessageWorker.SendDailyPredictions(ctx, receivers)
		if err != nil {
			return
		}
		defer s.Logger.Info("method send daily prediction request handled")
	}()
	return &pb.DailyPredictionsResponse{ResponseStatus: true}, nil
}
