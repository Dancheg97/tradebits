package main

import (
	"context"
	"fmt"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/user"
)

func (s *server) UserCreate(
	ctx context.Context,
	in *pb.UserCreateRequest,
) (*pb.Response, error) {
	senderAdress := calc.Hash(in.PublicKey)
	fmt.Println("craeting new user with name", in.PublicName)
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
