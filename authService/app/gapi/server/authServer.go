package server

import (
	"context"
	pb "github.com/Markuysa/astroMSA/apiGateway/app/protobuf/gen"
	"github.com/Markuysa/astroMSA/authService/app/internal/config"
	"github.com/Markuysa/astroMSA/authService/app/internal/database"
	"github.com/Markuysa/astroMSA/authService/app/internal/helpers/protobuf"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"go.uber.org/zap"
)

// Server struct - structure of gRPC server of the service
type Server struct {
	pb.UnimplementedAuthServiceServer
	Config  *config.Config
	UsersDB *database.UsersDB
	Port    string
	Logger  *zap.Logger
}

// NewServer method creates an object of the server
func NewServer(config *config.Config, usersDB *database.UsersDB, port string, logger *zap.Logger) *Server {
	return &Server{
		Config:  config,
		UsersDB: usersDB,
		Port:    port,
		Logger:  logger,
	}
}

// GetUser method handles the /api/v1/users/get endpoint
// and returns a protobuf user message
func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	s.Logger.Info("get user request received", zap.String("eMail", req.GetEMail()))
	user, err := s.UsersDB.Get(ctx, req.GetEMail())
	if err != nil {
		return nil, err
	}
	return protobuf.ConvertUserToPbResponse(user), nil
}

// AddUser method handles the /api/v1/users/add endpoint
// and adds some user, then returns the status code in AddUserResponse message
func (s *Server) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	s.Logger.Info("add user request received",
		zap.String("Email", req.Email),
		zap.String("Name", req.Name),
		zap.Bool("Notifications", req.Notifications),
		zap.Any("BirthDate", req.BirthInfo),
	)
	user := protobuf.ConvertUserRequestToUserStruct(req)
	err := s.UsersDB.Add(ctx, user)
	if err != nil {
		return &pb.AddUserResponse{Status: false}, err
	}
	return &pb.AddUserResponse{Status: true}, nil
}
