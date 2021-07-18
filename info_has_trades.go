package main

import (
	"context"
	"sync_tree/market"
	"sync_tree/user"

	pb "sync_tree/api"
)

func (s *server) InfoHasTrades(
	ctx context.Context,
	in *pb.InfoHasTradesRequest,
) (*pb.Response, error) {
	user := user.Look(in.UserAdress)
	if user != nil {
		market := market.Look(in.MarketAdress)
		if market != nil {
			has := market.HasTrades(in.UserAdress)
			return &pb.Response{Passed: has}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}
