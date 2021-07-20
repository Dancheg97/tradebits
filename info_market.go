package main

import (
	"context"
	"errors"
	"fmt"
	pb "sync_tree/api"
	"sync_tree/market"
)

func (s *server) InfoMarket(
	ctx context.Context,
	in *pb.InfoMarketRequest,
) (*pb.InfoMarketResponse, error) {
	m := market.Look(in.Adress)
	if m != nil {
		fmt.Println(m.Name)
		for idx, buy := range m.Pool.Buys {
			fmt.Println("buy", idx, "offer", buy.Offer, "recieve", buy.Recieve)
		}
		for idx, sell := range m.Pool.Sells {
			fmt.Println("sell", idx, "offer", sell.Offer, "recieve", sell.Recieve)
		}
		return &pb.InfoMarketResponse{
			MesKey:  m.MesKey,
			Name:    m.Name,
			Img:     m.Img,
			Descr:   m.Descr,
			OpCount: m.OpCount,
		}, nil
	}
	return &pb.InfoMarketResponse{}, errors.New("market not found")
}
