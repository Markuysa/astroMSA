package client

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/Markuysa/astroMSA/apigw/app/protobuf/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	authServerPort    = "9091"
	messageServerPort = "9090"
)

//func newAuthClient() (pbAuth.AuthServiceClient, error) {
//	authServerConnection, err := grpc.Dial(authServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
//	defer authServerConnection.Close()
//	if err != nil {
//		return nil, errors.New(fmt.Sprintf("failed to connect: %v", err))
//	}
//	authClient := pbAuth.NewAuthServiceClient(authServerConnection)
//	return &authClient, nil
//}
//
//func newMsgClient() (pbMsg.MessageServiceClient, error) {
//	messageServerConnection, err := grpc.Dial(messageServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
//	defer messageServerConnection.Close()
//	if err != nil {
//		return nil, errors.New(fmt.Sprintf("failed to connect: %v", err))
//	}
//
//	messageClient := pbMsg.NewMessageServiceClient(messageServerConnection)
//	return messageClient, nil
//}

// SendDailyPredictions Method to send the predictions to users
// Connects to msgSender - gRPC server and sends a request to send
// the predictions
func SendDailyPredictions(ctx context.Context) error {
	//connecting to msg sender service
	messageServerConnection, err := grpc.Dial(messageServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer messageServerConnection.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to connect: %v", err))
	}
	msgClient := pb.NewMessageServiceClient(messageServerConnection)
	//connecting to auth server
	authServerConnection, err := grpc.Dial(authServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer authServerConnection.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to connect: %v", err))
	}
	authClient := pb.NewAuthServiceClient(authServerConnection)

	//fetch the receivers of prediction
	receivers, err := authClient.GetUsersWithAllowedNotifications(ctx, &pb.NotificationsRequest{})
	if err != nil {
		return errors.New(fmt.Sprintf("failed to fetch receivers: %v", err))
	}

	_, err = msgClient.SendDailyPredictions(ctx, &pb.DailyPredictionsRequest{Day: "today", Receivers: receivers.Receivers})
	if err != nil {
		return err
	}
	return nil
}
