package client

import (
	"context"
	"errors"
	"fmt"
	pbAstro "github.com/Markuysa/astroMSA/astroService/app/protobuf/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	astroServerPort = "9091"
	authServerPort  = "9093"
)

// FetchPrediction method is needed to get a prediction
// from astroService using messageService gRPC client
func FetchPrediction(ctx context.Context, day string, sign string) (*pbAstro.PredictionResponse, error) {
	astroServiceConnection, err := grpc.Dial(astroServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer astroServiceConnection.Close()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to connect: %v", err))
	}
	astroClient := pbAstro.NewAstrologyServiceClient(astroServiceConnection)

	prediction, err := astroClient.GetPrediction(ctx, &pbAstro.PredictionRequest{Day: day, Sign: sign})
	if err != nil {
		return nil, err
	}
	return prediction, nil
}
