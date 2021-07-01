package main

import (
	"context"
	"fmt"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/user"
)

func (s *server) UserSend(
	ctx context.Context,
	in *pb.UserSendRequest,
) (*pb.Response, error) {
	fmt.Println("sending money, of amount: ", in.SendAmount)
	senderAdress := calc.Hash(in.PublicKey)
	amountBytes := calc.NumberToBytes(in.SendAmount)
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
		if sender != nil {
			defer sender.Save()
			if sender.Balance >= in.SendAmount {
				reciever := user.Get(in.RecieverAdress)
				if reciever != nil {
					defer reciever.Save()
					sender.Balance = sender.Balance - in.SendAmount
					reciever.Balance = reciever.Balance + in.SendAmount
					return &pb.Response{Passed: true}, nil
				}
			}
		}
	}
	return &pb.Response{Passed: false}, nil
}
