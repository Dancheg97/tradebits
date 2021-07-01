package main

import (
	"context"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/market"
	"sync_tree/user"
)

func (s *server) UserBuy(
	ctx context.Context,
	in *pb.UserBuyRequest,
) (*pb.Response, error) {
	buyerAdress := calc.Hash(in.PublicKey)
	buyer := user.Get(buyerAdress)
	if buyer != nil {
		if buyer.Balance >= in.Offer {
			offerBytes := calc.NumberToBytes(in.Offer)
			recieveBytes := calc.NumberToBytes(in.Recieve)
			concMessage := [][]byte{
				in.PublicKey,
				in.Adress,
				recieveBytes,
				offerBytes,
			}
			signErr := calc.Verify(concMessage, in.PublicKey, in.Sign)
			if signErr == nil {
				curMarket := market.Get(in.Adress)
				if curMarket != nil {
					trade := market.Trade{
						Adress:  buyerAdress,
						IsSell:  false,
						Offer:   in.Offer,
						Recieve: in.Recieve,
					}
					curMarket.OperateTrade(trade)
					return &pb.Response{Passed: true}, nil
				}
			}
		}
	}
	return &pb.Response{Passed: false}, nil
}
