package main

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/market"
	"sync_tree/trade"
	"sync_tree/user"
)

func (s *server) Create(
	ctx context.Context,
	in *pb.UserRequests_Create,
) (*pb.Response, error) {
	senderAdress := calc.Hash(in.PublicKey)
	signError := calc.Verify(
		[][]byte{
			in.PublicKey,
			in.MesssageKey,
			[]byte(in.PublicName),
		},
		in.PublicKey,
		in.Sign,
	)
	if signError != nil {
		fmt.Sprintln("[UserCreate] - Sign error")
		return nil, errors.New("user create error")

	}
	create_err := user.Create(
		senderAdress,
		in.MesssageKey,
		in.PublicName,
	)
	if create_err != nil {
		fmt.Sprintln("[UserCreate] - Create error")
		return nil, errors.New("create error")

	}
	fmt.Sprintln("[UserCreate] - User created")
	return &pb.Response{}, nil
}

func (s *server) Update(
	ctx context.Context,
	in *pb.UserRequests_Update,
) (*pb.Response, error) {
	senderAdress := calc.Hash(in.PublicKey)
	signError := calc.Verify(
		[][]byte{
			in.PublicKey,
			in.MesssageKey,
			[]byte(in.PublicName),
		},
		in.PublicKey,
		in.Sign,
	)
	if signError != nil {
		fmt.Sprintln("[UserUpdate] - Sign error")
		return nil, errors.New("sign check error")
	}
	user := user.Get(senderAdress)
	if user == nil {
		fmt.Sprintln("[UserUpdate] - User not found")
		return nil, errors.New("user not found error")
	}
	user.PublicName = in.PublicName
	user.Save()
	fmt.Sprintln("[UserUpdate] - User info updated: ", user.PublicName)
	return &pb.Response{}, nil
}

func (s *infoServer) UserSendMessage(
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
		u := user.Get(senderAdress)
		if u != nil {
			u.PutUserMessage(in.Adress, in.Message)
			u.Save()
			fmt.Sprintln("[UserSendMessage] - Message sent: ", u.PublicName)
			return &pb.Response{Passed: true}, nil
		}
		fmt.Sprintln("[UserSendMessage] - User not found error")
		return &pb.Response{Passed: false}, errors.New("user not found")
	}
	fmt.Sprintln("[UserSendMessage] - Sign error")
	return &pb.Response{Passed: false}, errors.New("sign error")
}

func (s *infoServer) UserSend(
	ctx context.Context,
	in *pb.UserSendRequest,
) (*pb.Response, error) {
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
	if !reflect.DeepEqual(senderAdress, in.RecieverAdress) {
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
						fmt.Sprintln("[UserSendMessage] - Message sent: ", sender.PublicName)
						return &pb.Response{Passed: true}, nil
					}
				}
			}
		}
	}
	fmt.Sprintln("[UserSendMessage] - User send error")
	return &pb.Response{Passed: false}, errors.New("send error")
}

func (s *infoServer) UserSell(
	ctx context.Context,
	in *pb.UserSellRequest,
) (*pb.Response, error) {
	sellerAdress := calc.Hash(in.PublicKey)
	seller := user.Get(sellerAdress)
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
				if !curMarket.HasTrades(sellerAdress) {
					trade := trade.Sell{
						Offer:   in.Offer,
						Recieve: in.Recieve,
					}
					attachedUsr := seller.AttachSell(&trade, in.Adress)
					if attachedUsr {
						attachedMkt := curMarket.AttachSell(&trade)
						if attachedMkt {
							fmt.Sprintln("[UserSell] - Sell order complete: ", seller.PublicName)
							return &pb.Response{Passed: true}, nil
						}
					}
				}
			}
		}
	}
	fmt.Sprintln("[UserSell] - User sell error")
	return &pb.Response{Passed: false}, errors.New("sell error")
}

func (s *infoServer) Buy(
	ctx context.Context,
	in *pb.UserRequests_Buy,
) (*pb.Response, error) {
	buyerAdress := calc.Hash(in.PublicKey)
	buyer := user.Get(buyerAdress)
	if buyer != nil {
		defer buyer.Save()
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
				if !curMarket.HasTrades(buyerAdress) {
					trade := trade.Buy{
						Offer:   in.Offer,
						Recieve: in.Offer,
					}
					attachedUsr := buyer.AttachBuy(&trade)
					if attachedUsr {
						attachedMkt := curMarket.AttachBuy(&trade)
						if attachedMkt {
							fmt.Sprintln("[UserBuy] - Buy order complete: ", buyer.PublicName)
							return &pb.Response{Passed: true}, nil
						}
					}
				}
			}
		}
	}
	fmt.Sprintln("[UserBuy] - User buy error")
	return &pb.Response{Passed: false}, errors.New("buy error")
}

func (s *infoServer) UserCancelTrade(
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
			fmt.Sprintln("[UserCancelTrade] - Trade canceled")
			return &pb.Response{Passed: true}, nil
		}
	}
	fmt.Sprintln("[UserCancelTrade] - Error canceling trade")
	return &pb.Response{Passed: false}, errors.New("cancel trade eeror")
}
