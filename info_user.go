package main

import (
	"context"
	"fmt"
	"sync_tree/user"

	pb "sync_tree/api"
)

func (s *server) InfoUser(
	ctx context.Context,
	in *pb.InfoUserRequest,
) (*pb.InfoUserResponse, error) {
	fmt.Println("giving information about", in.Adress)
	user := user.Look(in.Adress)
	if user == nil {
		return &pb.InfoUserResponse{
			PublicName: "====",
			Balance:    0,
		}, nil
	}
	adressesSlice := [][]byte{}
	balancesSlice := []uint64{}
	for strAdr, bal := range user.Markets {
		adressesSlice = append(adressesSlice, []byte(strAdr))
		balancesSlice = append(balancesSlice, bal)
	}
	return &pb.InfoUserResponse{
		PublicName:     user.PublicName,
		Balance:        user.Balance,
		MesKey:         user.MesKey,
		MarketAdresses: adressesSlice,
		MarketBalances: balancesSlice,
	}, nil
}
