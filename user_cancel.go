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
	concMes := [][]byte{
		in.PublicKey,
		in.MarketAdress,
	}
	verifyErr := calc.Verify(concMes, in.PublicKey, in.Sign)
	if verifyErr == nil {
		userAdress := calc.Hash(in.PublicKey)
		m := market.Get(in.MarketAdress)
		if m != nil {
			defer m.Save()
			canceled := m.CancelTrades(userAdress)
			fmt.Println(canceled)
			return &pb.Response{Passed: canceled}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}
