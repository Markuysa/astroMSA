package gapi

import (
	"authService/app/internal/config"
	"authService/app/internal/database"
	db "authService/app/internal/database"
	"authService/app/internal/helpers/protobuf"
	"authService/app/protobuf/pb"
	"context"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
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
	log.Println(s.UsersDB.GetUsersEmailsWithAllowedNotifications(ctx))
	user, err := s.UsersDB.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{User: &pb.User{
		Email:     user.Email,
		Sign:      user.Sign,
		Name:      user.Name,
		ID:        uint64(int64(user.ID)),
		BirthInfo: protobuf.DateToProtobuf(user.BirthInfo),
		CreatedAt: timestamppb.New(user.CreatedAt),
	}}, nil
}
func (s *Server) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {

}
func NewServerGateway() (*Server, error) {
	ctx := context.Background()
	usersDatabase := db.New(ctx)
	conf, err := config.New()
	if err != nil {
		return nil, nil
	}
	return NewServer(conf, usersDatabase, ":9090"), nil
}
