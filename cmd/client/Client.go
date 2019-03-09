package main

import (
	"fmt"
	v1 "google-speech/pkg/api/v1"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	h := v1.NewHealthClient(conn)

	fmt.Println(h)
	healthCheck, err := h.HealthCheck(context.Background(), &v1.HealthMessage{})

	fmt.Println(err)
	log.Printf("Response from health : %s \n", healthCheck.Health)

	tts := v1.NewTextToSpeechClient(conn)
	textTospch, err := tts.GenerateSpeech(context.Background(), &v1.TextToSpeechMessage{Data: "Convert This"})

	log.Printf("Response from Text to Speech : %s \n", textTospch.Response)
}
