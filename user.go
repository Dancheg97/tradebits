package main

import (
	"context"

	"sync_tree/calc"
	"sync_tree/lock"
	"sync_tree/user"

	pb "sync_tree/api"
)

func (s *server) NewUser(ctx context.Context, in *pb.UserCreateRequest) (*pb.Response, error) {
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
