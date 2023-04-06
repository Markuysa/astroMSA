package main

import (
	"context"
	"github.com/Markuysa/astroMSA/authService/app/gapi/server"
	"github.com/Markuysa/astroMSA/authService/app/internal/config"
	db "github.com/Markuysa/astroMSA/authService/app/internal/database"
	"github.com/Markuysa/astroMSA/authService/app/protobuf/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// runGRPC method starts a gRPC server of the service
func runGRPC(db *db.UsersDB, config *config.Config, port string) {
	grpcServer := grpc.NewServer()
	log.Print(port)
	authServer := server.NewServer(config, db, port)
	pb.RegisterAuthServiceServer(grpcServer, authServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", authServer.Port)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Starting grpc server at:", authServer.Port+" port")

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
func main() {

	//Change: port to be defined from docker cmpse
	ctx := context.Background()
	usersDatabase := db.New(ctx)
	config2, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	runGRPC(usersDatabase, config2, ":9093")

}
