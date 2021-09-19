package main

import (
	"context"
	"errors"
	"sync_tree/market"
	"sync_tree/search"
	"sync_tree/user"

	pb "sync_tree/api"
)

func (s *server) InfoHasTrades(
	ctx context.Context,
	in *pb.InfoHasTradesRequest,
) (*pb.Response, error) {
	// fmt.Println("info has trades call")
	user := user.Look(in.UserAdress)
	if user != nil {
		market := market.Look(in.MarketAdress)
		if market != nil {
			has := market.HasTrades(in.UserAdress)
			return &pb.Response{Passed: has}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}

func (s *server) InfoMarket(
	ctx context.Context,
	in *pb.InfoMarketRequest,
) (*pb.InfoMarketResponse, error) {
	// fmt.Println("info market call")
	m := market.Look(in.Adress)
	if m != nil {
		buys := []uint64{}
		for _, buy := range m.Pool.Buys {
			buys = append(buys, buy.Offer)
			buys = append(buys, buy.Recieve)
			if len(buys) == 10 {
				break
			}
		}
		sells := []uint64{}
		for _, sell := range m.Pool.Sells {
			sells = append(sells, sell.Offer)
			sells = append(sells, sell.Recieve)
			if len(sells) == 10 {
				break
			}
		}
		return &pb.InfoMarketResponse{
			MesKey:      m.MesKey,
			Name:        m.Name,
			Img:         m.Img,
			Descr:       m.Descr,
			OpCount:     m.OpCount,
			Buys:        buys,
			Sells:       sells,
			ActiveBuys:  uint64(len(buys) / 2),
			ActiveSells: uint64(len(sells) / 2),
			InputFee:    m.InputFee,
			OutputFee:   m.OutputFee,
			WorkTime:    m.WorkTime,
		}, nil
	}
	return &pb.InfoMarketResponse{}, errors.New("market not found")
}

func (s *server) InfoSearch(
	ctx context.Context,
	in *pb.InfoSearchRequest,
) (*pb.InfoSearchResponse, error) {
	// fmt.Println("user made a search request on: ", in.Info)
	results := search.Search(in.Info)
	if len(results) > 30 {
		results = results[0:30]
	}
	return &pb.InfoSearchResponse{ConcMarkets: results}, nil
}

func (s *server) InfoUser(
	ctx context.Context,
	in *pb.InfoUserRequest,
) (*pb.InfoUserResponse, error) {
	// fmt.Println("giving information about", in.Adress)
	user := user.Look(in.Adress)
	// fmt.Println(user)
	if user == nil {
		return &pb.InfoUserResponse{
			PublicName: "====",
			Balance:    0,
		}, nil
	}
	adressesSlice := [][]byte{}
	balancesSlice := []uint64{}
	for strAdr, bal := range user.Balances {
		adressesSlice = append(adressesSlice, []byte(strAdr))
		balancesSlice = append(balancesSlice, bal)
		// fmt.Println(adressesSlice)
		// fmt.Println(balancesSlice)
	}
	return &pb.InfoUserResponse{
		PublicName:     user.PublicName,
		Balance:        user.Balance,
		MesKey:         user.MesKey,
		MarketAdresses: adressesSlice,
		MarketBalances: balancesSlice,
	}, nil
}

func (s *server) InfoMessages(
	ctx context.Context,
	in *pb.InfoMessagesRequest,
) (*pb.Messages, error) {
	// fmt.Println("info messages request", in.UserAdress)
	usr := user.Look(in.UserAdress)
	msgs := usr.GetMessages(in.MarketAdress)
	return &pb.Messages{
		Messages: msgs,
	}, nil
}
