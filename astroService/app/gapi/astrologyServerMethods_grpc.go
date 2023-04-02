package gapi

import (
	"context"
	"github.com/Markuysa/astroMSA/astroService/app/internal/helpers"
	pb "github.com/Markuysa/astroMSA/astroService/app/protobuf/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func (s *Server) GetPrediction(ctx context.Context, req *pb.PredictionRequest) (*pb.PredictionResponse, error) {
	sign, day := req.Sign, req.Day
	s.logger.Info("Incoming get prediction request", zap.String("sign", sign), zap.String("day", day))
	predition, err := s.astrologyWorker.FetchPrediction(ctx, sign, day)
	if err != nil {
		return nil, err
	}
	return helpers.PredictionToPb(predition), nil
}

func (s *Server) SendPredictions(ctx context.Context, req *pb.SendPredictionsRequest) (*pb.SendPredictionsResponse, error) {
	conn, err := grpc.Dial("localhost:9004", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	return nil, nil
}
