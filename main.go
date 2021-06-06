package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	server := grpc.NewServer()
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to server grpc server over 9000: %v", err)
	}
}
