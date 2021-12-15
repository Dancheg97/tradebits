package market

import (
	"orb/calc"
	"orb/database"
	"orb/trade"
	"orb/user"
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
	database.TestRM([]byte(dummyName))
	database.TestRM(dummyAdress)
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
	database.TestRM([]byte(marketName))
	database.TestRM(marketAdress)
	database.TestRM([]byte(userName))
	database.TestRM(userAdress)
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
	database.TestRM([]byte(marketName))
	database.TestRM(marketAdress)
	database.TestRM([]byte(userName))
	database.TestRM(userAdress)
}
