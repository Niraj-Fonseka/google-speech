package cmd

import (
	"context"
	"google-speech/pkg/protocol/grpc"
	v1 "google-speech/pkg/service/v1"
)

func RunServer() error {
	ctx := context.Background()

	APIHealth := v1.NewHealthServiceServer()
	APITTs := v1.NewTextToSpeechServiceServer()

	return grpc.RunServer(ctx, APIHealth, APITTs, "8080")

}
