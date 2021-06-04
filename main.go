package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "sync_tree/api"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedSyncTreeServer
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
