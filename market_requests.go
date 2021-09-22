package main

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/market"
	"sync_tree/user"
)

func (s *server) Spawn(
	ctx context.Context,
	in *pb.MarketRequests_Create,
) (*pb.Response, error) {
	concatedMessage := [][]byte{
		in.PublicKey,
		in.MesssageKey,
		[]byte(in.Name),
		[]byte(in.Img),
		[]byte(in.Descr),
	}
	checkErr := calc.Verify(concatedMessage, in.PublicKey, in.Sign)
	if checkErr != nil {
		fmt.Println("[MarketCreate] - Error sign fail")
		return nil, errors.New("sign fail")
	}
	adress := calc.Hash(in.PublicKey)
	craeteErr := market.Create(
		adress,
		in.Name,
		in.MesssageKey,
		in.Descr,
		in.Img,
		in.InputFee,
		in.OutputFee,
		in.WorkTime,
		in.Delimiter,
	)
	if craeteErr != nil {
		fmt.Println("[MarketCreate] - Error create error: ", craeteErr)
		return nil, craeteErr
	}
	fmt.Println("[MarketCreate] - Market created, name: ", in.Name)
	return &pb.Response{}, nil
}

func (s *server) Refresh(
	ctx context.Context,
	in *pb.MarketRequests_Update,
) (*pb.Response, error) {
	concatedMessage := [][]byte{
		in.PublicKey,
		in.MesssageKey,
		[]byte(in.Name),
		[]byte(in.Img),
		[]byte(in.Descr),
	}
	checkErr := calc.Verify(concatedMessage, in.PublicKey, in.Sign)
	if checkErr != nil {
		fmt.Println("[MarketUpdate] - Sign error")
		return nil, errors.New("sign error")
	}
	adress := calc.Hash(in.PublicKey)
	mkt := market.Get(adress)
	if mkt == nil {
		fmt.Println("[MarketUpdate] - Market not found error")
		return nil, errors.New("sign error")
	}
	mkt.Name = in.Name
	mkt.MesKey = in.MesssageKey
	mkt.Descr = in.Descr
	mkt.Img = in.Img
	mkt.InputFee = in.InputFee
	mkt.OutputFee = in.OutputFee
	mkt.WorkTime = in.WorkTime
	mkt.Save()
	fmt.Println("[MarketUpdate] - Market info updated")
	return &pb.Response{}, nil
}

func (s *server) Deposit(
	ctx context.Context,
	in *pb.MarketRequests_Deposit,
) (*pb.Response, error) {
	amBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amBytes, uint64(in.Amount))
	concatedMessage := [][]byte{
		in.PublicKey,
		in.UserAdress,
		amBytes,
	}
	checkErr := calc.Verify(concatedMessage, in.PublicKey, in.Sign)
	if checkErr == nil {
		fmt.Println("[MarketDeposit] - Sign error")
		return nil, errors.New("sign error")
	}
	adress := calc.Hash(in.PublicKey)
	u := user.Get(in.UserAdress)
	if u == nil {
		fmt.Println("[MarketDeposit] - User not found error")
		return nil, errors.New("user not found error")
	}
	strAdr := string(adress)
	u.Balances[strAdr] = u.Balances[strAdr] + in.Amount
	fmt.Println("[MarketDeposit] - Deposit verified: ", u.PublicName)
	return &pb.Response{}, nil
}

func (s *server) Reply(
	ctx context.Context,
	in *pb.MarketRequests_Reply,
) (*pb.Response, error) {
	concMes := [][]byte{
		in.PublicKey,
		in.Adress,
		[]byte(in.Message),
	}
	signCheckErr := calc.Verify(concMes, in.PublicKey, in.Sign)
	if signCheckErr == nil {
		fmt.Println("[MarketSendMessage] - Sign error")
		return nil, errors.New("sign error")
	}
	senderAdress := calc.Hash(in.PublicKey)
	u := user.Get(senderAdress)
	if u == nil {
		fmt.Println("[MarketSendMessage] - User not found error")
		return nil, errors.New("sign error")
	}
	u.PutMarketMessage(in.Adress, in.Message)
	u.Save()
	fmt.Println("[MarketSendMessage] - Message sent", u.PublicName)
	return &pb.Response{}, nil
}

func (s *server) Withdrawal(
	ctx context.Context,
	in *pb.MarketRequests_Withdrawal,
) (*pb.Response, error) {
	amBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amBytes, uint64(in.Amount))
	concatedMessage := [][]byte{
		in.PublicKey,
		in.UserAdress,
		amBytes,
	}
	checkErr := calc.Verify(concatedMessage, in.PublicKey, in.Sign)
	if checkErr != nil {
		fmt.Println("[MarketWithdrawal] - Sign error")
		return nil, errors.New("sign error")
	}
	adress := calc.Hash(in.PublicKey)
	usr := user.Get(in.UserAdress)
	if usr == nil {
		fmt.Println("[MarketWithdrawal] - User not found error")
		return nil, errors.New("user not found")
	}
	defer usr.Save()
	strAdr := string(adress)
	if in.Amount > usr.Balances[strAdr] {
		fmt.Println("[MarketWithdrawal] - Withdrawal balance error")
		return nil, errors.New("bakance error")
	}
	usr.Balances[strAdr] = usr.Balances[strAdr] - in.Amount
	fmt.Println("[MarketWithdrawal] - Withdrawal accepted")
	return &pb.Response{}, nil
}
