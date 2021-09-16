package main

import (
	"context"
	"encoding/binary"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/market"
	"sync_tree/user"
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
				in.InputFee,
				in.OutputFee,
				in.WorkTime,
			)
			if craeteErr == nil {
				return &pb.Response{Passed: true}, nil
			}
		}
	}
	return &pb.Response{Passed: false}, nil
}

func (s *server) MarketDeposit(
	ctx context.Context,
	in *pb.MarketDepositRequest,
) (*pb.Response, error) {
	//fmt.Println("Operation market deposit for user: ", in.UserAdress)
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
			return &pb.Response{Passed: true}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
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
			return &pb.Response{Passed: true}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}

func (s *server) MarketUpdate(
	ctx context.Context,
	in *pb.MarketUpdateRequest,
) (*pb.Response, error) {
	//fmt.Println("got request to update market, with name: ", in.Name)
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
			return &pb.Response{Passed: true}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}

func (s *server) MarketWithdrawal(
	ctx context.Context,
	in *pb.MarketWithdrawalRequest,
) (*pb.Response, error) {
	//fmt.Println("Operation market withdrawal for user: ", in.UserAdress)
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
				return &pb.Response{Passed: true}, nil

			}
		}
	}
	return &pb.Response{Passed: false}, nil
}
