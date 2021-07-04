package main

import (
	"context"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/market"
)

func (s *server) MarketCreate(
	ctx context.Context,
	in *pb.MarketCreateRequest,
) (*pb.Response, error) {
	//fmt.Println("got request to craete new market, name", in.Name)
	concatedMessage := [][]byte{
		in.PublicKey,
		in.MesssageKey,
		[]byte(in.Name),
		[]byte(in.Img),
		[]byte(in.Descr),
	}
	checkErr := calc.Verify(concatedMessage, in.PublicKey, in.Sign)
	if checkErr == nil {
		if len(in.Name) < 15 {
			adress := calc.Hash(in.PublicKey)
			craeteErr := market.Create(
				adress,
				in.Name,
				in.MesssageKey,
				in.Descr,
				in.Img,
			)
			if craeteErr == nil {
				return &pb.Response{Passed: true}, nil
			}
		}
	}
	return &pb.Response{Passed: false}, nil
}
