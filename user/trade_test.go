package user

import (
	"orb/calc"
	"orb/data"
	"orb/trade"
	"reflect"
	"testing"
)

func TestAttachToLookedUser(t *testing.T) {
	var adress = calc.Rand()
	dummyName := string(calc.Rand()[0:8])
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Look(adress)
	buy := trade.Buy{}
	buyAttached := usr.AttachBuy(&buy)
	if buyAttached {
		t.Error("buy should not be attached to user, that can never be saved")
	}
	sell := trade.Sell{}
	sellAttached := usr.AttachSell(&sell, []byte{})
	if sellAttached {
		t.Error("sell trade should not be attached, cuz user can never be saved")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(adress)
}

func TestAttachTradesWithZeroOffer(t *testing.T) {
	var adress = calc.Rand()
	dummyName := string(calc.Rand()[0:8])
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Get(adress)
	buy := trade.Buy{
		Offer:   0,
		Recieve: 1000,
	}
	buyAttached := usr.AttachBuy(&buy)
	if buyAttached {
		t.Error("this buy should not be attached cuz hase zero offer")
	}
	sell := trade.Sell{
		Offer:   0,
		Recieve: 100,
	}
	sellAttached := usr.AttachSell(&sell, []byte{})
	if sellAttached {
		t.Error("this sell should never be attached cuz 0 offer")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(adress)
}

func TestAttachTradeWithBigBalance(t *testing.T) {
	var adress = calc.Rand()
	dummyName := string(calc.Rand()[0:8])
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Get(adress)
	buy := trade.Buy{
		Offer:   1000,
		Recieve: 1000,
	}
	buyAttached := usr.AttachBuy(&buy)
	if buyAttached {
		t.Error("this buy should not be attached cuz its over users balance")
	}
	sell := trade.Sell{
		Offer:   1000,
		Recieve: 1000,
	}
	usr.Balances["x"] = 0
	sellAttached := usr.AttachSell(&sell, []byte("x"))
	if sellAttached {
		t.Error("this sell should not be attached cuz its over users balance")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(adress)
}

func TestAttachNormalTrades(t *testing.T) {
	var adress = calc.Rand()
	dummyName := string(calc.Rand()[0:8])
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Get(adress)
	usr.Balance = 1000
	buy := trade.Buy{
		Offer:   1000,
		Recieve: 1000,
	}
	buyAttached := usr.AttachBuy(&buy)
	if !buyAttached {
		t.Error("this buy should be attached normally")
	}
	if !reflect.DeepEqual(buy.Adress, usr.adress) {
		t.Error("buy adress after bounding should be equal to users")
	}
	usr.Balances["x"] = 1000
	sell := trade.Sell{
		Offer:   1000,
		Recieve: 1000,
	}
	sellAttached := usr.AttachSell(&sell, []byte("x"))
	if !sellAttached {
		t.Error("this sell should be attached normally")
	}
	if !reflect.DeepEqual(sell.Adress, usr.adress) {
		t.Error("sell adress after bounding should be equal to users")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(adress)
}

func TestAttachSellNonExistingMarket(t *testing.T) {
	var adress = calc.Rand()
	dummyName := string(calc.Rand()[0:8])
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Get(adress)
	sell := trade.Sell{
		Offer:   1000,
		Recieve: 1000,
	}
	sellAttached := usr.AttachSell(&sell, []byte("x"))
	if sellAttached {
		t.Error("this sell should not be attached, cuz user dont have such market")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(adress)
}

func TestAttchBoundedTrades(t *testing.T) {
	var adress = calc.Rand()
	dummyName := string(calc.Rand()[0:8])
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Get(adress)
	usr.Balance = 100
	usr.Balances["x"] = 100
	buy := trade.Buy{
		Offer:   10,
		Recieve: 10,
		Adress:  []byte{0},
	}
	sell := trade.Sell{
		Offer:   10,
		Recieve: 10,
		Adress:  []byte{0},
	}
	buyAttached := usr.AttachBuy(&buy)
	sellAttached := usr.AttachSell(&sell, []byte("x"))
	if buyAttached || sellAttached {
		t.Error("those trades are already bounded and should not be attached")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(adress)
}
