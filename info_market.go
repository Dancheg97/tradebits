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
	fmt.Println("getting info about market: ", in.Adress)
	m := market.Look(in.Adress)
	if m != nil {
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
