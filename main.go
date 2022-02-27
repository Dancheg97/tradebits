package main

import (
	"log"
	"net"

	pb "setgraph/api"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedInfoServer
	pb.UnimplementedMarketServer
	pb.UnimplementedUserServer
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serv := grpc.NewServer()
	pb.RegisterInfoServer(serv, &server{})
	pb.RegisterUserServer(serv, &server{})
	pb.RegisterMarketServer(serv, &server{})
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
