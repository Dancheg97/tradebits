package main

import (
	"context"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/market"
)

func (s *server) MarketUpdate(
	ctx context.Context,
	in *pb.MarketUpdateRequest,
) (*pb.Response, error) {
	//fmt.Println("got request to update market, with name: ", in.Name)
	concatedMessage := [][]byte{
		in.PublicKey,
		in.MesssageKey,
		[]byte(in.Name),
		[]byte(in.Img),
		[]byte(in.Descr),
	}
	checkErr := calc.Verify(concatedMessage, in.PublicKey, in.Sign)
	if checkErr == nil {
		adress := calc.Hash(in.PublicKey)
		m := market.Get(adress)
		if m != nil {
			m.Name = in.Name
			m.MesKey = in.MesssageKey
			m.Descr = in.Descr
			m.Img = in.Img
			return &pb.Response{Passed: true}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}
