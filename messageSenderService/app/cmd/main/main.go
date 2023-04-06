package main

import (
	"github.com/Markuysa/astroMSA/messageSenderService/app/gapi/server"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/config"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/workers/cronWorker"
	"github.com/Markuysa/astroMSA/messageSenderService/app/internal/workers/messageSender"
	"github.com/Markuysa/astroMSA/messageSenderService/app/protobuf/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// runGRPC method runs a gRPC server on some port
func runGRPC(config2 *config.Config, cron *cronWorker.JobTicker, msgWorker *messageSender.MsgSenderWorker, port string) {
	grpcServer := grpc.NewServer()
	messageSenderServer := server.NewServer(config2, cron, msgWorker, port)
	pb.RegisterMessageServiceServer(grpcServer, messageSenderServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", messageSenderServer.Port)
	log.Println(listener)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Starting grpc server at:", messageSenderServer.Port+" port")

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
func main() {

	//ctx := context.Background()

	messageConfig, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	messageWorker := messageSender.Init(messageConfig)

	cron := cronWorker.Init()
	config2, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	runGRPC(config2, cron, messageWorker, ":9090")

}
