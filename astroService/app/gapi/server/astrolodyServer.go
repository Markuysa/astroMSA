package server

import (
	pb "github.com/Markuysa/astroMSA/apigw/app/protobuf/gen"
	"github.com/Markuysa/astroMSA/astroService/app/pkg/workers/astroWorker"
	"go.uber.org/zap"
)

// Server structure of the gRPC server
type Server struct {
	pb.UnimplementedAstrologyServiceServer
	Port            string
	astrologyWorker *astroWorker.AstroWorker
	logger          *zap.Logger
}

// NewServer creates an object of the server
func NewServer(worker *astroWorker.AstroWorker, logger *zap.Logger) *Server {
	return &Server{
		Port:            ":9091",
		astrologyWorker: worker,
		logger:          logger,
	}
}
