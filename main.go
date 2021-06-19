package main

import (
	"log"
	"net"
	"context"

	"sync_tree/calc"
	"sync_tree/lock"
	"sync_tree/user"

	"google.golang.org/grpc"
	pb "sync_tree/api"
)

const ()

type server struct {
	pb.UnimplementedSyncTreeServer
}

func (s *server) UserCreate(
	ctx context.Context,
	in *pb.UserCreateRequest,
) (*pb.Response, error) {
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
			return &pb.Response{Passed: true}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}

func main() {
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
