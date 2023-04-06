package main

import (
	"github.com/Markuysa/astroMSA/astroService/app/gapi/server"
	"github.com/Markuysa/astroMSA/astroService/app/internal/logger"
	pb "github.com/Markuysa/astroMSA/astroService/app/protobuf/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net/http"

	"github.com/Markuysa/astroMSA/astroService/app/pkg/workers/astroWorker"

	//"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// runGRPC Method to start a gRPC server of the service
func runGRPC(worker *astroWorker.AstroWorker, logger *zap.Logger) {
	grpcServer := grpc.NewServer()
	astroServer := server.NewServer(worker, logger)
	pb.RegisterAstrologyServiceServer(grpcServer, astroServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", astroServer.Port)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Starting grpc server at:", zap.String("port", astroServer.Port))

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	//ctx := context.Background()
	//config, err := config2.New()
	//
	//if err != nil {
	//	log.Println("WARINni")
	//	//log.Fatal(err)
	//}
	logger, err := logger.InitLogger()
	if err != nil {
		log.Fatal(err)
	}
	astrologyWorker := astroWorker.Init(nil, &http.Client{})

	runGRPC(astrologyWorker, logger)
	if err != nil {
		return
	}

}
