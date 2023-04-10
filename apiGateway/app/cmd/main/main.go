package main

import (
	"context"
	"fmt"
	protobuf "github.com/Markuysa/astroMSA/apiGateway/app/protobuf/gen"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

var (
	messageService = ":9090"
	astroService   = ":9091"
	authService    = ":9093"
)

// HTTPProxy REST API gateway for incoming requests
// Connects to every microservice using gRPC and
// runs a http server running on some port
func HTTPProxy(proxyaddr string) {
	grpcGwMux := runtime.NewServeMux()

	// Connecting to auth service
	grpcAuthConn, err := grpc.Dial(
		authService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Filed to connect to User service", err)
	}
	defer grpcAuthConn.Close()
	err = protobuf.RegisterAuthServiceHandler(
		context.Background(),
		grpcGwMux,
		grpcAuthConn,
	)
	if err != nil {
		log.Fatalln("Filed to start HTTP server", err)
	}
	//Connecting to message service
	grpcMessagesConn, err := grpc.Dial(
		messageService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Filed to connect to Messages sender service", err)
	}
	defer grpcMessagesConn.Close()

	err = protobuf.RegisterMessageServiceHandler(
		context.Background(),
		grpcGwMux,
		grpcMessagesConn,
	)
	if err != nil {
		log.Fatalln("Filed to start HTTP server", err)
	}
	// Connecting to astro service
	grpcAstroConn, err := grpc.Dial(
		astroService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Filed to connect to Astro service", err)
	}
	defer grpcAstroConn.Close()

	err = protobuf.RegisterAstrologyServiceHandler(
		context.Background(),
		grpcGwMux,
		grpcAstroConn,
	)
	if err != nil {
		log.Fatalln("Filed to start HTTP server", err)
	}

	mux := http.NewServeMux()

	mux.Handle("/api/v1/", grpcGwMux)

	fmt.Println("starting HTTP server at " + proxyaddr)
	log.Fatal(http.ListenAndServe(proxyaddr, mux))
}
func main() {
	HTTPProxy(":8086")
}
