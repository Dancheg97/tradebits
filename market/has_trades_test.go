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
	userAdress := calc.Rand()
	userName := string(calc.Rand()[0:8])
	user.Create(
		userAdress,
		dummyMessageKey,
		userName,
	)
	user := user.Get(userAdress)
	trade := trade.Buy{
		Offer:   1,
		Recieve: 1,
		Adress:  userAdress,
	}
	user.AttachBuy(&trade)
	mkt.AttachBuy(&trade)
	hasTrades := mkt.HasTrades(userAdress)
	if !hasTrades {
		t.Error("There should be a trade for that user")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(dummyAdress)
}

func TestHasSell(t *testing.T) {
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
