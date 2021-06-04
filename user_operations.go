package main

import (
	"context"

	pb "sync_tree/api"
)

func (s *server) NewUser(ctx context.Context, in *pb.UserCreateRequest) (*pb.Response, error) {
	message := [][]byte{}
	
	return &pb.Response{Passed: true}, nil
}
