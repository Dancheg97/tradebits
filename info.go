package main

import (
	"context"
	"errors"
	"fmt"
	"orb/data"
	"orb/market"
	"orb/search"
	"orb/user"

	pb "orb/api"
)

func (s *server) HasTrades(
	ctx context.Context,
	in *pb.InfIn_UserMarketAdresses,
) (*pb.InfOut_Bool, error) {
	fmt.Println("[InfoHasTrades] - start")
	user := user.Look(in.UserAdress)
	if user == nil {
		fmt.Println("[InfoHasTrades] - user not found")
		return nil, errors.New("user not found")
	}
	market := market.Look(in.MarketAdress)
	if market == nil {
		fmt.Println("[InfoHasTrades] - market not found")
		return nil, errors.New("market not found")
	}
	hasTrades := market.HasTrades(in.UserAdress)
	fmt.Println("[InfoHasTrades] - has trades - ", hasTrades)
	return &pb.InfOut_Bool{Value: hasTrades}, nil
}

func (s *server) Market(
	ctx context.Context,
	in *pb.InfIn_Adress,
) (*pb.InfOut_MarketInfo, error) {
	fmt.Println("[InfoMarket] - start")
	mkt := market.Look(in.Adress)
	if mkt == nil {
		fmt.Println("[InfoMarket] - market not found")
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
	fmt.Println("[InfoMarket] - info abound market: ", mkt.Name)
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
	in *pb.InfIn_Text,
) (*pb.InfOut_Adresses, error) {
	results := search.Search(in.Text)
	fmt.Println("[InfoSearch] - search results len: ", len(results))
	return &pb.InfOut_Adresses{MarketAdresses: results}, nil
}

func (s *server) CheckName(
	ctx context.Context,
	in *pb.InfIn_Text,
) (*pb.InfOut_Bool, error) {
	rez := data.Check([]byte(in.Text))
	fmt.Println("[CheckName] - Name:", in.Text, ", rez:", rez)
	return &pb.InfOut_Bool{Value: rez}, nil
}

func (s *server) User(
	ctx context.Context,
	in *pb.InfIn_Adress,
) (*pb.InfOut_User, error) {
	fmt.Println("[InfoUser] - start")
	user := user.Look(in.Adress)
	if user == nil {
		fmt.Println("[InfoUser] - error user not found")
		return nil, errors.New("user not found")
	}
	adressesSlice := [][]byte{}
	balancesSlice := []uint64{}
	for strAdr, bal := range user.Balances {
		adressesSlice = append(adressesSlice, []byte(strAdr))
		balancesSlice = append(balancesSlice, bal)
	}
	fmt.Println("[InfoUser] - info about user: ", user.Name)
	return &pb.InfOut_User{
		PublicName:     user.Name,
		Balance:        user.Balance,
		MessageKey:     user.MesKey,
		MarketAdresses: adressesSlice,
		MarketBalances: balancesSlice,
	}, nil
}

func (s *server) Messages(
	ctx context.Context,
	in *pb.InfIn_UserMarketAdresses,
) (*pb.InfOut_Messages, error) {
	fmt.Println("[InfoMessages] - start")
	usr := user.Look(in.UserAdress)
	messages := usr.GetMessages(in.MarketAdress)
	fmt.Println("[InfoMessages] - giving cipher messages of some user")
	return &pb.InfOut_Messages{
		Messages: messages,
	}, nil
}
