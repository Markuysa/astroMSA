package server

import (
	"context"
	pb "github.com/Markuysa/astroMSA/apiGateway/app/protobuf/gen"
	"github.com/Markuysa/astroMSA/astroService/app/gapi/client"
	"github.com/Markuysa/astroMSA/astroService/app/internal/helpers"
	"go.uber.org/zap"
)

// GetPrediction method returns a prediction
// by sign and day (today, tomorrow...)
func (s *Server) GetPrediction(ctx context.Context, req *pb.PredictionRequest) (*pb.PredictionResponse, error) {
	sign, day := req.Sign, req.Day
	s.logger.Info("Incoming get prediction request", zap.String("sign", sign), zap.String("day", day))
	predition, err := s.astrologyWorker.FetchPrediction(ctx, sign, day)
	if err != nil {
		return nil, err
	}
	return helpers.PredictionToPb(predition), nil
}

// SendPredictions method runs a client method
// SendDailyPredictions to send a predictions through other
// services
func (s *Server) SendPredictions(ctx context.Context, req *pb.SendPredictionsRequest) (*pb.SendPredictionsResponse, error) {
	err := client.SendDailyPredictions(ctx)
	if err != nil {
		return &pb.SendPredictionsResponse{Status: false}, err
	}
	return &pb.SendPredictionsResponse{Status: true}, nil
}
