package gapi

import (
	pb "astroService/app/protobuf/pb"
	"context"
)

func (s *Server) GetPrediction(ctx context.Context, req *pb.PredictionRequest) (*pb.PredictionResponse, error) {
	return nil, nil
}
