package main

import (
	"context"
	"log"
	pb "sync_tree/api"
	"sync_tree/calc"
	"testing"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func TestUserCreate(t *testing.T) {
	// TEST
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
	r, _ := c.AddUser(
		ctx,
		&pb.UserCreateRequest{
			PublicKey: keys.PersPub,
			ImgLink:   "nan",
			MesKey:    keys.MesPub,
			Sign:      sign,
		},
	)
	if !r.Passed {
		t.Error("user have not been created")
	}
}
