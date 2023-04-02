package main

import (
	"authService/app/gapi"
	"authService/app/internal/config"
	db "authService/app/internal/database"
	"authService/app/protobuf/pb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func runGRPC(db *db.UsersDB, config *config.Config, port string) {
	grpcServer := grpc.NewServer()
	log.Print(port)
	authServer := gapi.NewServer(config, db, port)
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
	runGRPC(usersDatabase, config2, ":9004")

}
