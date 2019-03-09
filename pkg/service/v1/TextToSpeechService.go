package v1

import (
	v1 "google-speech/pkg/api/v1"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type TextToSpeechServiceServer struct {
}

func NewTextToSpeechServiceServer() v1.TextToSpeechServer {
	return &TextToSpeechServiceServer{}
}

func (s *TextToSpeechServiceServer) GenerateSpeech(ctx context.Context, in *v1.TextToSpeechMessage) (*v1.TextToSpeechMessage, error) {
	return &v1.TextToSpeechMessage{Response: "Audio Generated"}, nil
}
