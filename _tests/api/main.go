package main

import (
	"context"
	"log"
	"testing"
	"time"
	"sync_tree/calc"
	pb "sync_tree/api"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSyncTreeClient(conn)

	keys := calc.Gen()
	message := [][]byte{
		keys.PersPub,
		keys.MesPub,
		[]byte("nan"),
	}
	sign, _ := calc.Sign(message, keys.PersPriv)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddUser(
		ctx,
		&pb.UserCreateRequest{
			PublicKey: keys.PersPub,
			ImgLink:   "nan",
			MesKey:    keys.MesPriv,
			Sign:      sign,
		},
	)
	if !r.Passed {
		t.Error("user was not created")
	}
}
