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

func (s *server) MarketCreate(
	ctx context.Context,
	in *pb.MarketCreateRequest,
) (*pb.Response, error) {
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
		craeteErr := market.Create(
			adress,
			in.Name,
			in.MesssageKey,
			in.Descr,
			in.Img,
			in.InputFee,
			in.OutputFee,
			in.WorkTime,
		)
		if craeteErr == nil {
			fmt.Sprintln("[MarketCreate] - Market created, name: ", in.Name)
			return &pb.Response{Passed: true}, nil
		}
		fmt.Sprintln("[MarketCreate] - Error create error: ", craeteErr)
		return &pb.Response{Passed: false}, craeteErr
	}
	fmt.Sprintln("[MarketCreate] - Error sign fail")
	return &pb.Response{Passed: false}, errors.New("sign fail")
}

func (s *server) MarketDeposit(
	ctx context.Context,
	in *pb.MarketDepositRequest,
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
		adress := calc.Hash(in.PublicKey)
		u := user.Get(in.UserAdress)
		if u != nil {
			strAdr := string(adress)
			u.Balances[strAdr] = u.Balances[strAdr] + in.Amount
			fmt.Sprintln("[MarketDeposit] - Deposit verified: ", u.PublicName)
			return &pb.Response{Passed: true}, nil
		}
		fmt.Sprintln("[MarketDeposit] - User not found error")
		return &pb.Response{Passed: false}, errors.New("user not found error")
	}
	fmt.Sprintln("[MarketDeposit] - Sign error")
	return &pb.Response{Passed: false}, errors.New("sign error")
}

func (s *server) MarketSendMessage(
	ctx context.Context,
	in *pb.MarketSendMessageRequest,
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
			u.PutMarketMessage(in.Adress, in.Message)
			u.Save()
			fmt.Sprintln("[MarketSendMessage] - Message sent", u.PublicName)
			return &pb.Response{Passed: true}, nil
		}
		fmt.Sprintln("[MarketSendMessage] - User not found error")
		return &pb.Response{Passed: false}, errors.New("sign error")
	}
	fmt.Sprintln("[MarketSendMessage] - Sign error")
	return &pb.Response{Passed: false}, errors.New("sign error")
}

func (s *server) MarketUpdate(
	ctx context.Context,
	in *pb.MarketUpdateRequest,
) (*pb.Response, error) {
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
			m.InputFee = in.InputFee
			m.OutputFee = in.OutputFee
			m.WorkTime = in.WorkTime
			m.Save()
			fmt.Sprintln("[MarketUpdate] - Market info updated")
			return &pb.Response{Passed: true}, nil
		}
		fmt.Sprintln("[MarketUpdate] - Market not found error")
		return &pb.Response{Passed: false}, errors.New("sign error")
	}
	fmt.Sprintln("[MarketUpdate] - Sign error")
	return &pb.Response{Passed: false}, errors.New("sign error")
}

func (s *server) MarketWithdrawal(
	ctx context.Context,
	in *pb.MarketWithdrawalRequest,
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
		adress := calc.Hash(in.PublicKey)
		u := user.Get(in.UserAdress)
		if u != nil {
			strAdr := string(adress)
			if in.Amount < u.Balances[strAdr] {
				u.Balances[strAdr] = u.Balances[strAdr] - in.Amount
				fmt.Sprintln("[MarketWithdrawal] - Withdrawal accepted")
				return &pb.Response{Passed: true}, nil
			}
			fmt.Sprintln("[MarketWithdrawal] - Withdrawal balance error")
			return &pb.Response{Passed: false}, errors.New("bakance error")
		}
		fmt.Sprintln("[MarketWithdrawal] - User not found error")
		return &pb.Response{Passed: false}, errors.New("user not found")
	}
	fmt.Sprintln("[MarketWithdrawal] - Sign error")
	return &pb.Response{Passed: false}, errors.New("sign error")
}
