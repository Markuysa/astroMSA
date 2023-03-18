package main

import (
	"github.com/Markuysa/astroMSA/astroService/app/gapi"
	pb "github.com/Markuysa/astroMSA/astroService/app/protobuf/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	//"astroService/gapi"
	config2 "github.com/Markuysa/astroMSA/astroService/app/internal/config"
	"github.com/Markuysa/astroMSA/astroService/app/pkg/workers/astroWorker"
	//pb "astroService/protobuf"
	"github.com/valyala/fasthttp"
	//"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

func runGRPC(worker *astroWorker.AstroWorker) {
	grpcServer := grpc.NewServer()
	astroServer := gapi.NewServer(worker)
	pb.RegisterAstrologyServiceServer(grpcServer, astroServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", astroServer.Port)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Starting grpc server at:", astroServer.Port+" port")

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	//ctx := context.Background()
	config, err := config2.New()
	if err != nil {
		log.Fatal(err)
	}
	client := fasthttp.Client{ReadTimeout: 5 * time.Second}
	astrologyWorker := astroWorker.Init(config, &client)

	runGRPC(astrologyWorker)
	if err != nil {
		return
	}

}
