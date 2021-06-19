package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"sync_tree/calc"
	"sync_tree/lock"
	"sync_tree/user"

	pb "sync_tree/api"

	"google.golang.org/grpc"
)

const ()

type server struct {
	pb.UnimplementedSyncTreeServer
}

func (s *server) UserCreate(
	ctx context.Context,
	in *pb.UserCreateRequest,
) (*pb.Response, error) {
	fmt.Println("hello phone")
	senderAdress := calc.Hash(in.PublicKey)
	lock.Lock(senderAdress)
	defer lock.Unlock(senderAdress)
	check_err := calc.Verify(
		[][]byte{
			in.PublicKey,
			in.MesssageKey,
			[]byte(in.PublicName),
		},
		in.PublicKey,
		in.Sign,
	)
	if check_err == nil {
		create_err := user.Create(
			senderAdress,
			in.MesssageKey,
			in.PublicName,
		)
		if create_err == nil {
			fmt.Println("created")
			return &pb.Response{Passed: true}, nil
		}
	}
	fmt.Println("not created")
	return &pb.Response{Passed: false}, nil
}

func main() {
	fmt.Println("the game goes on")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSyncTreeServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
