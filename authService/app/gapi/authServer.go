package gapi

import (
	"apigw/app/services/pb"
	"authService/app/internal/config"
	"authService/app/internal/database"
	db "authService/app/internal/database"
	"context"
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
	return &pb.GetUserResponse{User: &pb.User{
		Email: user.Email,
		Sign:  user.Sign,
		Name:  user.Name,
		ID:    uint64(int64(user.ID)),
		//BirthInfo: user.BirthInfo,
		//CreatedAt: user.CreatedAt,
	}}, nil
}

func NewServerGateway() (*Server, error) {
	ctx := context.Background()
	usersDatabase := db.New(ctx)
	config, err := config.New()
	if err != nil {
		return nil, nil
	}
	return NewServer(config, usersDatabase, ":9090"), nil
}
