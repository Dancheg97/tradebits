package main

import (
	"context"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/user"
)

func (s *server) MarketSendMessage(
	ctx context.Context,
	in *pb.MarketSendMessageRequest,
) (*pb.Response, error) {
	concMes := [][]byte{
		in.PublicKey,
		in.Adress,
		[]byte(in.Message),
	}
	signCheckErr := calc.Verify(concMes, in.PublicKey, in.Sign)
	if signCheckErr == nil {
		senderAdress := calc.Hash(in.PublicKey)
		u := user.Get(in.Adress)
		if u != nil {
			u.PutMessage(senderAdress, in.Message)
			return &pb.Response{Passed: true}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}
