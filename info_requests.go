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

func (s *server) HasTrades(
	ctx context.Context,
	in *pb.InfIn_UserMarketAdresses,
) (*pb.Response, error) {
	user := user.Look(in.UserAdress)
	if user == nil {
		fmt.Sprintln("[InfoHasTrades] - user not found")
		return nil, errors.New("user not found")
	}
	market := market.Look(in.MarketAdress)
	if market == nil {
		fmt.Sprintln("[InfoHasTrades] - market not found")
		return nil, errors.New("market not found")
	}
	hasTrades := market.HasTrades(in.UserAdress)
	fmt.Sprintln("[InfoHasTrades] - has trades - ", hasTrades)
	return &pb.Response{}, nil

}

func (s *server) Market(
	ctx context.Context,
	in *pb.InfIn_Adress,
) (*pb.InfOut_MarketInfo, error) {
	mkt := market.Look(in.Adress)
	if mkt == nil {
		fmt.Sprintln("[InfoMarket] - market not found")
		return nil, errors.New("market not found")

	}
	buys := []*pb.InfOut_Trade{}
	for _, buy := range mkt.Pool.Buys {
		buys = append(buys, &pb.InfOut_Trade{
			Offer:   buy.Offer,
			Recieve: buy.Recieve,
		})
		if len(buys) == 10 {
			break
		}
	}
	sells := []*pb.InfOut_Trade{}
	for _, sell := range mkt.Pool.Sells {
		sells = append(sells, &pb.InfOut_Trade{
			Offer:   sell.Offer,
			Recieve: sell.Recieve,
		})
		if len(sells) == 10 {
			break
		}
	}
	fmt.Sprintln("[InfoMarket] - info abound market: ", mkt.Name)
	return &pb.InfOut_MarketInfo{
		MessageKey:     mkt.MesKey,
		Name:           mkt.Name,
		ImageLink:      mkt.Img,
		Description:    mkt.Descr,
		OperationCount: mkt.OpCount,
		Buys:           buys,
		Sells:          sells,
		ActiveBuys:     uint64(len(buys) / 2),
		ActiveSells:    uint64(len(sells) / 2),
		InputFee:       mkt.InputFee,
		OutputFee:      mkt.OutputFee,
		WorkTime:       mkt.WorkTime,
		Delimiter:      mkt.Delimiter,
	}, nil
}

func (s *server) Search(
	ctx context.Context,
	in *pb.InfIn_SearchText,
) (*pb.InfOut_Adresses, error) {
	results := search.Search(in.Text)
	fmt.Sprintln("[InfoSearch] - search results len: ", len(results))
	return &pb.InfOut_Adresses{MarketAdresses: results}, nil
}

func (s *infoServer) InfoUser(
	ctx context.Context,
	in *pb.InfoUserRequest,
) (*pb.InfoUserResponse, error) {
	user := user.Look(in.Adress)
	if user == nil {
		fmt.Sprintln("[InfoUser] - error user not found")
		return &pb.InfoUserResponse{}, errors.New("user not found")
	}
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

func (s *infoServer) InfoMessages(
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
