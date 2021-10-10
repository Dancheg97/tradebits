package market

import (
	"sync_tree/calc"
	"sync_tree/data"
	"sync_tree/trade"
	"sync_tree/user"
	"testing"
	"time"
)

func TestCancelSell(t *testing.T) {
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
	usr.Balances[string(marketAdress)] = 100
	sell := trade.Sell{
		Offer:   100,
		Recieve: 100,
	}
	usr.AttachSell(&sell, marketAdress)
	mkt.AttachSell(&sell)
	mkt.CancelTrades(userAdress)
	usr.Save()
	mkt.Save()
	time.Sleep(time.Millisecond * 300)
	checkMkt := Look(marketAdress)
	checkUsr := user.Look(userAdress)
	if len(checkMkt.Pool.Sells) != 0 {
		t.Error("market sell length should be 0")
	}
	if checkUsr.Balances[string(marketAdress)] != 100 {
		t.Error("user mkt balance should be 100")
	}
	data.TestRM([]byte(userName))
	data.TestRM(userAdress)
	data.TestRM([]byte(marketName))
	data.TestRM(marketAdress)
}

func TestCancelBuy(t *testing.T) {
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
	usr.Balance = 100
	buy := trade.Buy{
		Offer:   100,
		Recieve: 100,
	}
	usr.AttachBuy(&buy)
	mkt.AttachBuy(&buy)
	mkt.CancelTrades(userAdress)
	usr.Save()
	mkt.Save()
	time.Sleep(time.Millisecond * 300)
	checkMkt := Look(marketAdress)
	checkUsr := user.Look(userAdress)
	if len(checkMkt.Pool.Sells) != 0 {
		t.Error("market sell length should be 0")
	}
	if checkUsr.Balance != 100 {
		t.Error("user mkt balance should be 100")
	}
	data.TestRM([]byte(userName))
	data.TestRM(userAdress)
	data.TestRM([]byte(marketName))
	data.TestRM(marketAdress)
}
