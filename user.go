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
	fmt.Println("[UserCreate] - start")
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
		fmt.Println("[UserCreate] - Sign error")
		return nil, errors.New("user create error")

	}
	create_err := user.Create(
		senderAdress,
		in.MesssageKey,
		in.PublicName,
	)
	if create_err != nil {
		fmt.Println("[UserCreate] - Create error", create_err)
		return nil, create_err

	}
	fmt.Println("[UserCreate] - User created")
	return &pb.Response{}, nil
}

func (s *server) Send(
	ctx context.Context,
	in *pb.UserRequests_Send,
) (*pb.Response, error) {
	fmt.Println("[UserSend] - start")
	senderAdress := calc.Hash(in.PublicKey)
	if reflect.DeepEqual(senderAdress, in.RecieverAdress) {
		fmt.Println("[UserSend] - Reciever is sender")
		return nil, errors.New("reciever is sender")
	}
	sender := user.Get(senderAdress)
	if sender == nil {
		fmt.Println("[UserSend] - Sender dont exist")
		return nil, errors.New("sender dont exist")
	}
	defer sender.Save()
	reciever := user.Get(in.RecieverAdress)
	if reciever == nil {
		fmt.Println("[UserSend] - Reciever dont exits")
		return nil, errors.New("reciever dont exist")
	}
	defer reciever.Save()
	if sender.Balance < in.SendAmount {
		fmt.Println("[UserSend] - Not enough balance")
		return nil, errors.New("not enough balance")
	}
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
	if signError != nil {
		fmt.Println("[UserSendMessage] - Sign error")
		return nil, errors.New("sign error")
	}
	sender.Balance = sender.Balance - in.SendAmount
	reciever.Balance = reciever.Balance + in.SendAmount
	fmt.Println("[UserSendMessage] - Message sent: ", sender.Name)
	return &pb.Response{}, nil
}

func (s *server) Message(
	ctx context.Context,
	in *pb.UserRequests_Message,
) (*pb.Response, error) {
	fmt.Println("[Message] - start")
	concMes := [][]byte{
		in.PublicKey,
		in.Adress,
		in.Message,
	}
	signCheckErr := calc.Verify(concMes, in.PublicKey, in.Sign)
	if signCheckErr != nil {
		fmt.Println("[UserMessage] - Sign error")
		return nil, errors.New("sign error")
	}
	senderAdress := calc.Hash(in.PublicKey)
	u := user.Get(senderAdress)
	if u == nil {
		fmt.Println("[UserMessage] - User not found error")
		return nil, errors.New("user not found")
	}
	defer u.Save()
	u.PutMessage(in.Adress, in.Message)
	fmt.Println("[UserMessage] - Message sent: ", u.Name)
	return &pb.Response{}, nil
}

func (s *server) Buy(
	ctx context.Context,
	in *pb.UserRequests_Trade,
) (*pb.Response, error) {
	fmt.Println("[UserBuy] - start")
	buyerAdress := calc.Hash(in.PublicKey)
	buyer := user.Get(buyerAdress)
	if buyer == nil {
		fmt.Println("[UserBuy] - Buyer dont exists")
		return nil, errors.New("buyer dont exist")
	}
	defer buyer.Save()
	curMarket := market.Get(in.Adress)
	if curMarket == nil {
		fmt.Println("[UserBuy] - Market dont exists")
		return nil, errors.New("market dont exist")
	}
	defer curMarket.Save()
	trade := trade.Buy{
		Offer:   in.Offer,
		Recieve: in.Recieve,
	}
	concMessage := [][]byte{
		in.PublicKey,
		in.Adress,
		calc.NumberToBytes(in.Recieve),
		calc.NumberToBytes(in.Offer),
	}
	signErr := calc.Verify(concMessage, in.PublicKey, in.Sign)
	if signErr != nil {
		fmt.Println("[UserBuy] - Sign check fail")
		return nil, errors.New("sign check fail")
	}
	if curMarket.HasTrades(buyerAdress) {
		fmt.Println("[UserBuy] - Has active trades")
		return nil, errors.New("has active trades")
	}
	attachedToUser := buyer.AttachBuy(&trade)
	if !attachedToUser {
		fmt.Println("[UserBuy] - Trade user attach fail")
		return nil, errors.New("trade user attach fail")
	}
	attachedToMarket := curMarket.AttachBuy(&trade)
	if !attachedToMarket {
		fmt.Println("[UserBuy] - Trade market attach fail")
		return nil, errors.New("trade market attach fail")
	}
	fmt.Println("[UserBuy] - Buy order complete: ", buyer.Name)
	return &pb.Response{}, nil
}

func (s *server) Sell(
	ctx context.Context,
	in *pb.UserRequests_Trade,
) (*pb.Response, error) {
	fmt.Println("[UserSell] - start")
	sellerAdress := calc.Hash(in.PublicKey)
	seller := user.Get(sellerAdress)
	if seller == nil {
		fmt.Println("[UserSell] - Seller dont exists")
		return nil, errors.New("seller dont exist")
	}
	defer seller.Save()
	curMarket := market.Get(in.Adress)
	if curMarket == nil {
		fmt.Println("[UserSell] - Market dont exists")
		return nil, errors.New("market dont exist")
	}
	defer curMarket.Save()
	trade := trade.Sell{
		Offer:   in.Offer,
		Recieve: in.Recieve,
	}
	concMessage := [][]byte{
		in.PublicKey,
		in.Adress,
		calc.NumberToBytes(in.Recieve),
		calc.NumberToBytes(in.Offer),
	}
	signErr := calc.Verify(concMessage, in.PublicKey, in.Sign)
	if signErr != nil {
		fmt.Println("[UserSell] - Sign check fail")
		return nil, errors.New("sign check fail")
	}
	if curMarket.HasTrades(sellerAdress) {
		fmt.Println("[UserSell] - Has active trades")
		return nil, errors.New("has active trades")
	}
	attachedToUser := seller.AttachSell(&trade, in.Adress)
	if !attachedToUser {
		fmt.Println("[UserSell] - Trade user attach fail")
		return nil, errors.New("trade user attach fail")
	}
	attachedToMarket := curMarket.AttachSell(&trade)
	if !attachedToMarket {
		fmt.Println("[UserSell] - Trade market attach fail")
		return nil, errors.New("trade market attach fail")
	}
	fmt.Println("[UserSell] - Sell order complete: ", seller.Name)
	return &pb.Response{}, nil
}

func (s *server) CancelTrade(
	ctx context.Context,
	in *pb.UserRequests_CancelTrade,
) (*pb.Response, error) {
	fmt.Println("[UserCancelTrade] - start")
	mkt := market.Get(in.MarketAdress)
	if mkt == nil {
		fmt.Println("[CancelTrade] - No such market")
		return nil, errors.New("no such market")
	}
	defer mkt.Save()
	userAdress := calc.Hash(in.PublicKey)
	concMes := [][]byte{
		in.PublicKey,
		in.MarketAdress,
	}
	verifyErr := calc.Verify(concMes, in.PublicKey, in.Sign)
	if verifyErr != nil {
		fmt.Println("[CancelTrade] - Sign error")
		return nil, errors.New("sign error")
	}
	mkt.CancelTrades(userAdress)
	fmt.Println("[UserCancelTrade] - Trade canceled successfully")
	return &pb.Response{}, nil
}
