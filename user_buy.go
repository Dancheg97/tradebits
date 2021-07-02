package main

import (
	"context"
	"fmt"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/market"
	"sync_tree/user"
)

func (s *server) UserBuy(
	ctx context.Context,
	in *pb.UserBuyRequest,
) (*pb.Response, error) {
	fmt.Println("buy: [offer / recieve] [", in.Offer, "/", in.Recieve, "]")
	buyerAdress := calc.Hash(in.PublicKey)
	buyer := user.Get(buyerAdress)
	if buyer != nil {
		fmt.Println("ok1")
		if buyer.Balance >= in.Offer {
			fmt.Println("ok2")
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
				fmt.Println("ok3")
				curMarket := market.Get(in.Adress)
				if curMarket != nil {
					fmt.Println("ok4")
					buyer.Balance = buyer.Balance - in.Offer
					buyer.Save()
					trade := market.Trade{
						Adress:  buyerAdress,
						IsSell:  false,
						Offer:   in.Offer,
						Recieve: in.Recieve,
					}
					curMarket.OperateTrade(trade)
					curMarket.Save()
					return &pb.Response{Passed: true}, nil
				}
			}
		}
	}
	return &pb.Response{Passed: false}, nil
}
