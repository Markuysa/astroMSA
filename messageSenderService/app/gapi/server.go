package gapi

import (
	"messageSenderService/app/internal/config"
	"messageSenderService/app/internal/workers/cronWorker"
	"messageSenderService/app/internal/workers/messageSender"
	"messageSenderService/app/protobuf/pb"
)

type Server struct {
	pb.UnimplementedMessageServiceServer
	MessageWorker *messageSender.MsgSenderWorker
	CronWorker    *cronWorker.JobTicker
	Config        *config.Config
	Port          string
}

func NewServer(config *config.Config, cronWorker *cronWorker.JobTicker, msgWorker *messageSender.MsgSenderWorker, port string) *Server {
	return &Server{
		Config:        config,
		MessageWorker: msgWorker,
		CronWorker:    cronWorker,
		Port:          port,
	}
}
