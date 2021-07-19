package main

import (
	"context"
	"fmt"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/market"
	"sync_tree/trade"
	"sync_tree/user"
)

func (s *server) UserSell(
	ctx context.Context,
	in *pb.UserSellRequest,
) (*pb.Response, error) {
	fmt.Println("sell offer / recieve: ", in.Offer, "/", in.Recieve)
	buyerAdress := calc.Hash(in.PublicKey)
	seller := user.Get(buyerAdress)
	if seller != nil {
		defer seller.Save()
		curMarket := market.Get(in.Adress)
		if curMarket != nil {
			defer curMarket.Save()
			concMessage := [][]byte{
				in.PublicKey,
				in.Adress,
				calc.NumberToBytes(in.Recieve),
				calc.NumberToBytes(in.Offer),
			}
			signErr := calc.Verify(concMessage, in.PublicKey, in.Sign)
			if signErr == nil {
				trade := trade.Sell{
					Offer:   in.Offer,
					Recieve: in.Recieve,
				}
				attachedUsr := seller.AttachSell(&trade, in.Adress)
				if attachedUsr {
					attachedMkt := curMarket.AttachSell(&trade)
					if attachedMkt {
						return &pb.Response{Passed: true}, nil
					}
				}
			}
		}
	}
	return &pb.Response{Passed: false}, nil
}
