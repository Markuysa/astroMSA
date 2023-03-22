package gapi

import (
	"context"
	"github.com/Markuysa/astroMSA/astroService/app/internal/helpers"
	pb "github.com/Markuysa/astroMSA/astroService/app/protobuf/pb"
)

func (s *Server) GetPrediction(ctx context.Context, req *pb.PredictionRequest) (*pb.PredictionResponse, error) {

	sign, day := req.Sign, req.Sign
	predition, err := s.astrologyWorker.FetchPrediction(ctx, sign, day)
	if err != nil {
		return nil, err
	}
	return helpers.PredictionToPb(predition), nil
}
