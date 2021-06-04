package main

import (
	"log"
	"net"

	"context"
	"fmt"

	"sync_tree/calc"
	"sync_tree/lock"
	"sync_tree/user"

	"google.golang.org/grpc"
	pb "sync_tree/api"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedSyncTreeServer
}

func (s *server) AddUser(ctx context.Context, in *pb.UserCreateRequest) (*pb.Response, error) {
	fmt.Println("recieved request")
	senderAdress := calc.Hash(in.PublicKey)
	lock.Lock(senderAdress)
	defer lock.Unlock(senderAdress)
	check_err := calc.Verify(
		[][]byte{
			in.PublicKey,
			in.MesKey,
			[]byte(in.ImgLink),
		},
		in.PublicKey,
		in.Sign,
	)
	if check_err == nil {
		create_err := user.Create(senderAdress, in.MesKey, in.ImgLink)
		if create_err == nil {
			return &pb.Response{Passed: true}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSyncTreeServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
