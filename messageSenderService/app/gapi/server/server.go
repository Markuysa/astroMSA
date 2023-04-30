package server

import (
	pb "github.com/Markuysa/astroMSA/apiGateway/app/protobuf/gen"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/config"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/workers/cronWorker"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/workers/messageSender"
	"go.uber.org/zap"
)

// Server is a structure of gRPC server
// of the service
type Server struct {
	pb.UnimplementedMessageServiceServer
	MessageWorker *messageSender.MsgSenderWorker
	CronWorker    *cronWorker.JobTicker
	Config        *config.Config
	Port          string
	Logger        *zap.Logger
}

// NewServer method returns an object of messageSenderService
func NewServer(config *config.Config,
	cronWorker *cronWorker.JobTicker,
	msgWorker *messageSender.MsgSenderWorker,
	port string,
	logger *zap.Logger) *Server {

	return &Server{
		Config:        config,
		MessageWorker: msgWorker,
		CronWorker:    cronWorker,
		Port:          port,
		Logger:        logger,
	}
}
