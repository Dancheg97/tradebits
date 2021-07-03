package main

import (
	"context"
	"fmt"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/market"
)

func (s *server) UserCancelTrade(
	ctx context.Context,
	in *pb.UserCancelTradeRequest,
) (*pb.Response, error) {
	fmt.Println("cancel:", in.MarketAdress)
	concMes := [][]byte{
		in.PublicKey,
		in.MarketAdress,
	}
	verifyErr := calc.Verify(concMes, in.PublicKey, in.Sign)
	if verifyErr == nil {
		userAdress := calc.Hash(in.PublicKey)
		m := market.Get(in.MarketAdress)
		if m != nil {
			canceled := m.CancelTrades(userAdress)
			return &pb.Response{Passed: canceled}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}
