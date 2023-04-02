package gapi

import (
	"context"
	"github.com/Markuysa/astroMSA/authService/app/internal/config"
	"github.com/Markuysa/astroMSA/authService/app/internal/database"
	"github.com/Markuysa/astroMSA/authService/app/internal/helpers/protobuf"
	"github.com/Markuysa/astroMSA/authService/app/protobuf/pb"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
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
	user, err := s.UsersDB.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return protobuf.ConvertUserToPbResponse(user), nil
}
func (s *Server) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	user := protobuf.ConvertUserRequestToUserStruct(req)

	err := s.UsersDB.Add(ctx, user)
	if err != nil {
		return &pb.AddUserResponse{Status: false}, err
	}
	return &pb.AddUserResponse{Status: true}, nil
}
