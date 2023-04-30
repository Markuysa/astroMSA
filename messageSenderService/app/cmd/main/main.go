package main

import (
	pb "github.com/Markuysa/astroMSA/apiGateway/app/protobuf/gen"
	"github.com/Markuysa/astroMSA/messageSenderService/app/gapi/server"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/config"
	logger2 "github.com/Markuysa/astroMSA/messageSenderService/app/internal/logger"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/workers/cronWorker"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/workers/messageSender"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// runGRPC method runs a gRPC server on some port
func runGRPC(messageSenderServer *server.Server) {
	grpcServer := grpc.NewServer()
	pb.RegisterMessageServiceServer(grpcServer, messageSenderServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", messageSenderServer.Port)
	log.Println(listener)
	if err != nil {
		log.Fatal(err)
	}
	messageSenderServer.Logger.Info("starting grpc server", zap.String("port", messageSenderServer.Port))

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}

var ServerPort = ":9091"

func main() {

	messageConfig, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	logger, err := logger2.InitLogger()
	if err != nil {
		log.Fatal(err)
	}
	messageWorker := messageSender.Init(messageConfig)

	cron := cronWorker.Init()
	config2, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	messageSenderServer := server.NewServer(
		config2,
		cron,
		messageWorker,
		ServerPort,
		logger,
	)
	runGRPC(messageSenderServer)

}
