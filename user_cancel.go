package main

import (
	"context"
	"fmt"
	pb "sync_tree/api"
	"sync_tree/calc"
)

func (s *server) UserCancelTrade(
	ctx context.Context,
	in *pb.UserCancelTradeRequest,
) (*pb.Response, error) {
	fmt.Println("cancel:", in.MarketAdress)
	userAdress := calc.Hash(in.PublicKey)
	concMes := [][]byte{
		in.PublicKey,
		in.MarketAdress,
	}
	verifyErr := calc.Verify(concMes, in.PublicKey, in.Sign)
	if verifyErr == nil {
		
	}
	return &pb.Response{Passed: false}, nil
}
