package main

import (
	"fmt"
	"google-speech/api"
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

	h := api.NewHealthClient(conn)

	fmt.Println(h)
	healthCheck, err := h.HealthCheck(context.Background(), &api.HealthMessage{})

	fmt.Println(err)
	log.Printf("Response from health : %s \n", healthCheck.Health)

}
