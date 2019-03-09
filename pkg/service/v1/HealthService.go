package v1

import (
	v1 "google-speech/pkg/api/v1"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type HealthServiceServer struct {
}

func NewHealthServiceServer() v1.HealthServer {
	return &HealthServiceServer{}
}

func (s *HealthServiceServer) HealthCheck(ctx context.Context, in *v1.HealthMessage) (*v1.HealthMessage, error) {
	return &v1.HealthMessage{Health: "Health is good ! ", Status: true}, nil
}
