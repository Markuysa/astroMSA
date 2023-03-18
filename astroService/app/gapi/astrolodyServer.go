package gapi

import (
	"github.com/Markuysa/astroMSA/astroService/app/pkg/workers/astroWorker"
	pb "github.com/Markuysa/astroMSA/astroService/app/protobuf/pb"
)

type Server struct {
	pb.UnimplementedAstrologyServiceServer
	Port            string
	astrologyWorker *astroWorker.AstroWorker
}

func NewServer(worker *astroWorker.AstroWorker) *Server {
	return &Server{
		Port:            ":9090",
		astrologyWorker: worker,
	}
}
