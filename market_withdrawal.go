package main

import (
	"context"
	"encoding/binary"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/user"
)

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
			if in.Amount < u.Markets[strAdr] {
				u.Markets[strAdr] = u.Markets[strAdr] - in.Amount
				return &pb.Response{Passed: true}, nil

			}
		}
	}
	return &pb.Response{Passed: false}, nil
}
