package main

import (
	"context"
	"fmt"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/market"
	"sync_tree/user"
)

func (s *server) UserSell(
	ctx context.Context,
	in *pb.UserSellRequest,
) (*pb.Response, error) {
	fmt.Println("sell offer / recieve: ", in.Offer, "/", in.Recieve)
	buyerAdress := calc.Hash(in.PublicKey)
	buyer := user.Get(buyerAdress)
	if buyer != nil {
		mktAdr := string(in.Adress)
		if buyer.Markets[mktAdr] >= in.Offer {
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
					buyer.Markets[mktAdr] = buyer.Markets[mktAdr] - in.Offer
					fmt.Println(buyer)
					buyer.Save()
					trade := market.Trade{
						Adress:  buyerAdress,
						IsSell:  true,
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
