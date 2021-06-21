package main

import (
	"context"
	"fmt"
	"sync_tree/user"

	pb "sync_tree/api"
)

func (s *server) UserInfo(
	ctx context.Context,
	in *pb.UserInfoRequest,
) (*pb.UserInfoResponse, error) {
	fmt.Println("giving information about", in.Adress)
	user := user.Look(in.Adress)
	if user == nil {
		return &pb.UserInfoResponse{
			PublicName: "====",
			Balance:    0,
		}, nil
	}
	return &pb.UserInfoResponse{
		PublicName: user.PublicName,
		Balance:    user.Balance,
	}, nil
}
