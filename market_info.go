package main

import (
	"context"
	"errors"
	pb "sync_tree/api"
	"sync_tree/market"
)

func (s *server) MarketInfo(
	ctx context.Context,
	in *pb.MarketInfoRequest,
) (*pb.MarketInfoResponse, error) {
	m := market.Look(in.Adress)
	if m == nil {
		return &pb.MarketInfoResponse{}, errors.New("market not found")
	}
	return &pb.MarketInfoResponse{
		MesssageKey: m.MesKey,
		Name:        m.Name,
		Img:         m.Img,
		Descr:       m.Descr,
		OpCount:     m.OpCount,
	}, nil
}
