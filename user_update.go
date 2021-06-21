package main

import (
	"context"
	"fmt"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/user"
)

func (s *server) UserUpdate(
	ctx context.Context,
	in *pb.UserUpdateRequest,
) (*pb.Response, error) {
	senderAdress := calc.Hash(in.PublicKey)
	fmt.Println("updating user name", in.PublicName)
	signError := calc.Verify(
		[][]byte{
			in.PublicKey,
			in.MesssageKey,
			[]byte(in.PublicName),
		},
		in.PublicKey,
		in.Sign,
	)
	if signError == nil {
		user := user.Get(senderAdress)
		if user != nil {
			user.PublicName = in.PublicName
			user.Save()
			return &pb.Response{Passed: true}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}
