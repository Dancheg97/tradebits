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

func (s *server) UserBuy(
	ctx context.Context,
	in *pb.UserBuyRequest,
) (*pb.Response, error) {
	fmt.Println("buy: [offer / recieve] [", in.Offer, "/", in.Recieve, "]")
	buyerAdress := calc.Hash(in.PublicKey)
	buyer := user.Get(buyerAdress)
	if buyer != nil {
		defer buyer.Save()
		curMarket := market.Get(in.Adress)
		if curMarket != nil {
			defer curMarket.Save()
			fmt.Println("passed")
			concMessage := [][]byte{
				in.PublicKey,
				in.Adress,
				calc.NumberToBytes(in.Recieve),
				calc.NumberToBytes(in.Offer),
			}
			signErr := calc.Verify(concMessage, in.PublicKey, in.Sign)
			if signErr == nil {
				trade := trade.Buy{
					Offer:   in.Offer,
					Recieve: in.Offer,
				}
				attachedUsr := buyer.AttachBuy(&trade)
				if attachedUsr {
					attachedMkt := curMarket.AttachBuy(&trade)
					if attachedMkt {
						return &pb.Response{Passed: true}, nil
					}
				}
			}
		}
	}
	return &pb.Response{Passed: false}, nil
}

func (s *server) UserCreate(
	ctx context.Context,
	in *pb.UserCreateRequest,
) (*pb.Response, error) {
	senderAdress := calc.Hash(in.PublicKey)
	fmt.Println("craeting new user with name", in.PublicName)
	signError := calc.Verify(
		[][]byte{
			in.PublicKey,
			in.MesssageKey,
			[]byte(in.PublicName),
		},
		in.PublicKey,
		in.Sign,
	)
	if signError == nil {
		create_err := user.Create(
			senderAdress,
			in.MesssageKey,
			in.PublicName,
		)
		if create_err == nil {
			return &pb.Response{Passed: true}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}

func (s *server) UserSendMessage(
	ctx context.Context,
	in *pb.UserSendMessageRequest,
) (*pb.Response, error) {
	concMes := [][]byte{
		in.PublicKey,
		in.Adress,
		[]byte(in.Message),
	}
	signCheckErr := calc.Verify(concMes, in.PublicKey, in.Sign)
	if signCheckErr == nil {
		senderAdress := calc.Hash(in.PublicKey)
		u := user.Get(in.Adress)
		if u != nil {
			u.PutUserMessage(senderAdress, in.Message)
			u.Save()
			return &pb.Response{Passed: true}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}

func (s *server) UserSend(
	ctx context.Context,
	in *pb.UserSendRequest,
) (*pb.Response, error) {
	fmt.Println("sending money, of amount: ", in.SendAmount)
	senderAdress := calc.Hash(in.PublicKey)
	amountBytes := calc.NumberToBytes(in.SendAmount)
	signError := calc.Verify(
		[][]byte{
			in.PublicKey,
			amountBytes,
			in.RecieverAdress,
		},
		in.PublicKey,
		in.Sign,
	)
	if signError == nil {
		sender := user.Get(senderAdress)
		if sender != nil {
			defer sender.Save()
			if sender.Balance >= in.SendAmount {
				reciever := user.Get(in.RecieverAdress)
				if reciever != nil {
					defer reciever.Save()
					sender.Balance = sender.Balance - in.SendAmount
					reciever.Balance = reciever.Balance + in.SendAmount
					return &pb.Response{Passed: true}, nil
				}
			}
		}
	}
	return &pb.Response{Passed: false}, nil
}

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

func (s *server) UserUpdate(
	ctx context.Context,
	in *pb.UserUpdateRequest,
) (*pb.Response, error) {
	senderAdress := calc.Hash(in.PublicKey)
	fmt.Println("updating user name", in.PublicName)
	signError := calc.Verify(
		[][]byte{
			in.PublicKey,
			in.MesssageKey,
			[]byte(in.PublicName),
		},
		in.PublicKey,
		in.Sign,
	)
	if signError == nil {
		user := user.Get(senderAdress)
		if user != nil {
			user.PublicName = in.PublicName
			user.Save()
			return &pb.Response{Passed: true}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}

func (s *server) UserCancelTrade(
	ctx context.Context,
	in *pb.UserCancelTradeRequest,
) (*pb.Response, error) {
	concMes := [][]byte{
		in.PublicKey,
		in.MarketAdress,
	}
	verifyErr := calc.Verify(concMes, in.PublicKey, in.Sign)
	if verifyErr == nil {
		userAdress := calc.Hash(in.PublicKey)
		m := market.Get(in.MarketAdress)
		if m != nil {
			defer m.Save()
			m.CancelTrades(userAdress)
			return &pb.Response{Passed: true}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}
