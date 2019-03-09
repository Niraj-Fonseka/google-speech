package cmd

import (
	"context"
	"google-speech/pkg/protocol/grpc"
	v1 "google-speech/pkg/service/v1"
)

func RunServer() error {
	ctx := context.Background()

	v1API := v1.NewHealthServiceServer()

	return grpc.RunServer(ctx, v1API, "8080")

}
