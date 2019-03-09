package api

import (
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

func (s *Server) HealthCheck(ctx context.Context, in *HealthMessage) (*HealthMessage, error) {
	return &HealthMessage{Health: "Health is good ! ", Status: true}, nil
}
