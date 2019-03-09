package main

import (
	"fmt"
	cmd "google-speech/pkg/cmd"
	"os"
)

// main start a gRPC server and waits for connection
func main() {

	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v \n", err)
		os.Exit(1)
	}
}
