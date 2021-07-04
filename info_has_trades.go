package main

import (
	"context"
	"reflect"
	"sync_tree/market"
	"sync_tree/user"

	pb "sync_tree/api"
)

func (s *server) InfoHasTrades(
	ctx context.Context,
	in *pb.InfoHasTradesRequest,
) (*pb.Response, error) {
	//fmt.Println("checking trades")
	user := user.Look(in.UserAdress)
	if user != nil {
		market := market.Look(in.MarketAdress)
		if market != nil {
			has := false
			for _, trade := range market.Buys {
				if reflect.DeepEqual(trade.Adress, in.UserAdress) {
					has = true
				}
			}
			for _, trade := range market.Sells {
				if reflect.DeepEqual(trade.Adress, in.UserAdress) {
					has = true
				}
			}
			return &pb.Response{Passed: has}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}
