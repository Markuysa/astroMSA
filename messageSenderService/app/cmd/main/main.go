package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"messageSenderService/app/gapi"
	"messageSenderService/app/internal/config"
	"messageSenderService/app/internal/workers/cronWorker"
	"messageSenderService/app/internal/workers/messageSender"
	"messageSenderService/app/protobuf/pb"
	"net"
)

func runGRPC(config2 *config.Config, cron *cronWorker.JobTicker, msgWorker *messageSender.MsgSenderWorker, port string) {
	grpcServer := grpc.NewServer()
	messageSenderServer := gapi.NewServer(config2, cron, msgWorker, port)
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

	runGRPC(config2, cron, messageWorker, ":9092")

}
