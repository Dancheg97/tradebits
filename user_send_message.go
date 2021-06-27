package main

import (
	"context"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/market"
)

func (s *server) UserSendMessage(
	ctx context.Context,
	in *pb.UserSendMessageRequest,
) (*pb.Response, error) {
	concMes := [][]byte{
		in.PublicKey,
		in.Adress,
		[]byte(in.Message),
	}
	signCheckErr := calc.Verify(concMes, in.PublicKey, in.Sign)
	if signCheckErr == nil {
		senderAdress := calc.Hash(in.PublicKey)
		m := market.Get(in.Adress)
		m.PutMessage(senderAdress, in.Message)
	}
	return &pb.Response{Passed: false}, nil
}
