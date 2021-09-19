package main

import (
	"context"
	"errors"
	"fmt"
	"sync_tree/market"
	"sync_tree/search"
	"sync_tree/user"

	pb "sync_tree/api"
)

func (s *server) InfoHasTrades(
	ctx context.Context,
	in *pb.InfoHasTradesRequest,
) (*pb.Response, error) {
	user := user.Look(in.UserAdress)
	if user != nil {
		market := market.Look(in.MarketAdress)
		if market != nil {
			hasTrades := market.HasTrades(in.UserAdress)
			fmt.Sprintln("[InfoHasTrades] - has trades")
			return &pb.Response{Passed: hasTrades}, nil
		}
		fmt.Sprintln("[InfoHasTrades] - market not found")
		return &pb.Response{Passed: false}, errors.New("market not found")
	}
	fmt.Sprintln("[InfoHasTrades] - user not found")
	return &pb.Response{Passed: false}, errors.New("user not found")
}

func (s *server) InfoMarket(
	ctx context.Context,
	in *pb.InfoMarketRequest,
) (*pb.InfoMarketResponse, error) {
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
		fmt.Sprintln("[InfoMarket] - info abound market: ", m.Name)
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
	fmt.Sprintln("[InfoMarket] - market not found")
	return &pb.InfoMarketResponse{}, errors.New("market not found")
}

func (s *server) InfoSearch(
	ctx context.Context,
	in *pb.InfoSearchRequest,
) (*pb.InfoSearchResponse, error) {
	results := search.Search(in.Info)
	if len(results) > 30 {
		results = results[0:30]
	}
	fmt.Sprintln("[InfoSearch] - search results len: ", len(results))
	return &pb.InfoSearchResponse{ConcMarkets: results}, nil
}

func (s *server) InfoUser(
	ctx context.Context,
	in *pb.InfoUserRequest,
) (*pb.InfoUserResponse, error) {
	user := user.Look(in.Adress)
	if user != nil {
		adressesSlice := [][]byte{}
		balancesSlice := []uint64{}
		for strAdr, bal := range user.Balances {
			adressesSlice = append(adressesSlice, []byte(strAdr))
			balancesSlice = append(balancesSlice, bal)
		}
		fmt.Sprintln("[InfoUser] - info about user: ", user.PublicName)
		return &pb.InfoUserResponse{
			PublicName:     user.PublicName,
			Balance:        user.Balance,
			MesKey:         user.MesKey,
			MarketAdresses: adressesSlice,
			MarketBalances: balancesSlice,
		}, nil
	}
	fmt.Sprintln("[InfoUser] - error user not found")
	return &pb.InfoUserResponse{}, errors.New("user not found")
}

func (s *server) InfoMessages(
	ctx context.Context,
	in *pb.InfoMessagesRequest,
) (*pb.Messages, error) {
	usr := user.Look(in.UserAdress)
	msgs := usr.GetMessages(in.MarketAdress)
	fmt.Sprintln("[InfoMessages] - giving cipher messages of some user")
	return &pb.Messages{
		Messages: msgs,
	}, nil
}
