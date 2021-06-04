package main

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "sync_tree/api"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func TestUserCreateRequest() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSyncTreeClient(conn)

	new_user = 

}
