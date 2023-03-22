package main

import (
	protobuf "apigw/app/services/pb"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

var (
	authService    = ":9000"
	messageService = ":9002"
	astroService   = ":9090"
)

func HTTPProxy(proxyaddr string) {
	grpcGwMux := runtime.NewServeMux()

	// Connecting to auth service
	grpcAuthConn, err := grpc.Dial(
		authService,
		//grpc.WithPerRPCCredentials(&reqData{}),
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
		//grpc.WithPerRPCCredentials(&reqData{}),
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
	// Connecting to auth service
	grpcAstroConn, err := grpc.Dial(
		astroService,
		//grpc.WithPerRPCCredentials(&reqData{}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Filed to connect to User service", err)
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

//	func runGatewayServer() {
//		authServer, err := gapi.NewServerGateway()
//		if err != nil {
//			log.Fatal(err)
//		}
//		grpcMux := runtime.NewServeMux()
//
//		ctx, cancel := context.WithCancel(context.Background())
//		defer cancel()
//		err = protobuf.RegisterAuthServiceHandlerServer(ctx, grpcMux, authServer)
//		if err != nil {
//			log.Fatal(err)
//		}
//		mux := http.NewServeMux()
//		mux.Handle("/", grpcMux)
//		listener, err := net.Listen("tcp", ":8080")
//		if err != nil {
//			log.Fatalln(err)
//		}
//
//		log.Println("starting http gateway")
//		err = http.Serve(listener, mux)
//		if err != nil {
//			log.Fatalln(err)
//		}
//	}
//
//	func runGrpcServer() {
//		server, err := gapi.NewServerGateway()
//		if err != nil {
//			log.Fatalln(err)
//		}
//
//		grpcServer := grpc.NewServer()
//		protobuf.RegisterAuthServiceServer(grpcServer, server)
//		reflection.Register(grpcServer)
//
//		listener, err := net.Listen("tcp", ":9096")
//		if err != nil {
//			log.Fatalln(err)
//		}
//
//		log.Println("starting grpc")
//		err = grpcServer.Serve(listener)
//		if err != nil {
//			log.Fatalln(err)
//		}
//
// }
func main() {
	HTTPProxy(":8080")
	//go runGrpcServer()
	//runGatewayServer()

}
