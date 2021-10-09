package market

import (
	"sync_tree/calc"
	"sync_tree/data"
	"sync_tree/trade"
	"sync_tree/user"
	"testing"
)

func TestHasNoTrades(t *testing.T) {
	dummyAdress := calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	Create(
		dummyAdress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	mkt := Get(dummyAdress)
	dummyUserAdress := calc.Rand()
	hasTrades := mkt.HasTrades(dummyUserAdress)
	if hasTrades {
		t.Error("new market should not have any trades")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(dummyAdress)
}

func TestHasBuy(t *testing.T) {
	marketAdress := calc.Rand()
	marketName := string(calc.Rand()[0:16])
	Create(
		marketAdress,
		marketName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	mkt := Get(marketAdress)
	userAdress := calc.Rand()
	userName := string(calc.Rand()[0:8])
	user.Create(
		userAdress,
		dummyMessageKey,
		userName,
	)
	usr := user.Get(userAdress)
	usr.Balance = 1
	trade := trade.Buy{
		Offer:   1,
		Recieve: 1,
		Adress:  userAdress,
	}
	usr.AttachBuy(&trade)
	mkt.AttachBuy(&trade)
	hasTrades := mkt.HasTrades(userAdress)
	if !hasTrades {
		t.Error("There should be a trade for that user")
	}
	data.TestRM([]byte(marketName))
	data.TestRM(marketAdress)
	data.TestRM([]byte(userName))
	data.TestRM(userAdress)
}

func TestHasSell(t *testing.T) {
	marketAdress := calc.Rand()
	marketName := string(calc.Rand()[0:16])
	Create(
		marketAdress,
		marketName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	mkt := Get(marketAdress)
	userAdress := calc.Rand()
	userName := string(calc.Rand()[0:8])
	user.Create(
		userAdress,
		dummyMessageKey,
		userName,
	)
	usr := user.Get(userAdress)
	usr.Balances[string(marketAdress)] = 1
	trade := trade.Sell{
		Offer:   1,
		Recieve: 1,
		Adress:  userAdress,
	}
	usr.AttachSell(&trade, marketAdress)
	mkt.AttachSell(&trade)
	hasTrades := mkt.HasTrades(userAdress)
	if !hasTrades {
		t.Error("There should be a trade for that user")
	}
	data.TestRM([]byte(marketName))
	data.TestRM(marketAdress)
	data.TestRM([]byte(userName))
	data.TestRM(userAdress)
}
