package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"sync_tree/calc"
	"sync_tree/lock"
	"sync_tree/user"

	pb "sync_tree/api"

	"google.golang.org/grpc"
)

const ()

type server struct {
	pb.UnimplementedSyncTreeServer
}

func (s *server) UserCreate(
	ctx context.Context,
	in *pb.UserCreateRequest,
) (*pb.Response, error) {
	fmt.Println("mes got")
	senderAdress := calc.Hash(in.PublicKey)
	lock.Lock(senderAdress)
	defer lock.Unlock(senderAdress)
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

func (s *server) UserChangePubName(
	ctx context.Context,
	in *pb.UserUpdateRequest,
) (*pb.Response, error) {
	fmt.Println("mes got")
	senderAdress := calc.Hash(in.PublicKey)
	lock.Lock(senderAdress)
	defer lock.Unlock(senderAdress)
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
		user.PublicName = in.PublicName
		user.Save()
		return &pb.Response{Passed: true}, nil
	}
	return &pb.Response{Passed: false}, nil
}

func (s *server) UserSendRequest(
	ctx context.Context,
	in *pb.UserSendRequest,
) (*pb.Response, error) {
	fmt.Println("mes got")
	senderAdress := calc.Hash(in.PublicKey)
	lock.Lock(senderAdress)
	defer lock.Unlock(senderAdress)
	amountBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountBytes, in.SendAmount)
	signError := calc.Verify(
		[][]byte{
			in.PublicKey,
			amountBytes,
			in.RecieverAdress,
		},
		in.PublicKey,
		in.Sign,
	)
	if signError == nil {
		sender := user.Get(senderAdress)
		if sender.Balance > in.SendAmount {
			reciever := user.Get(in.RecieverAdress)
			if reciever != nil {
				lock.Lock(in.RecieverAdress)
				defer lock.Unlock(in.RecieverAdress)
				sender.Balance = sender.Balance - in.SendAmount
				reciever.Balance = reciever.Balance + in.SendAmount
				sender.Save()
				reciever.Save()
				return &pb.Response{Passed: true}, nil
			}
		}
	}
	return &pb.Response{Passed: false}, nil
}

func main() {
	fmt.Println("the game goes on")
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
