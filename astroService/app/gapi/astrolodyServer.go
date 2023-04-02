package gapi

import (
	"github.com/Markuysa/astroMSA/astroService/app/pkg/workers/astroWorker"
	pb "github.com/Markuysa/astroMSA/astroService/app/protobuf/pb"
	"go.uber.org/zap"
)

type Server struct {
	pb.UnimplementedAstrologyServiceServer
	Port            string
	astrologyWorker *astroWorker.AstroWorker
	logger          *zap.Logger
}

func NewServer(worker *astroWorker.AstroWorker, logger *zap.Logger) *Server {
	return &Server{
		Port:            ":9003",
		astrologyWorker: worker,
		logger:          logger,
	}
}
