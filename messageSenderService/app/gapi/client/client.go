package client

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/Markuysa/astroMSA/apiGateway/app/protobuf/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	astroServerPort = "9090"
	authServerPort  = "9092"
)

// FetchPrediction method is needed to get a prediction
// from astroService using messageService gRPC client
func FetchPrediction(ctx context.Context, day string, sign string) (*pb.PredictionResponse, error) {
	astroServiceConnection, err := grpc.Dial(astroServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer astroServiceConnection.Close()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to connect: %v", err))
	}
	astroClient := pb.NewAstrologyServiceClient(astroServiceConnection)
	log.Println(day, sign)
	prediction, err := astroClient.GetPrediction(ctx, &pb.PredictionRequest{Day: day, Sign: sign})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(prediction)
	return prediction, nil
}
