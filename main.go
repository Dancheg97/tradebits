package main

import (
	"fmt"
	"log"
	"net"

	pb "sync_tree/api"
	"sync_tree/api_test_preparation"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedInfoServer
	pb.UnimplementedMarketServer
	pb.UnimplementedUserServer
}

func main() {
<<<<<<< HEAD
	createNewUsers()
	createStartMarket()
=======
	api_test_preparation.CreateNewUsers()
	api_test_preparation.CreateNewMarkets()
>>>>>>> ebae813c220497fb4b90f6557ae6034f04d3999a
	fmt.Println("Server has started! port: 8080 open")
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
