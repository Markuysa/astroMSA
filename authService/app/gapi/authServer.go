package gapi

import (
	"authService/app/internal/config"
	"authService/app/internal/database"
	"authService/app/protobuf/pb"
	"context"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	Config  *config.Config
	UsersDB *database.UsersDB
	Port    string
}

func NewServer(config *config.Config, usersDB *database.UsersDB, port string) *Server {
	return &Server{
		Config:  config,
		UsersDB: usersDB,
		Port:    port,
	}

}
func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return nil, nil
}
