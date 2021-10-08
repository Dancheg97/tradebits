package market

import (
	"sync_tree/calc"
	"sync_tree/data"
	"sync_tree/trade"
	"sync_tree/user"
	"testing"
	"time"
)

var dummyMessageKey = []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2}
var dummyImageLink = "test.imagelink/thereisnoimagebythislink"
var dummyDescription = `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrbled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets ctaining Lorem Ipsum pages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`
var dummyInputFee = uint64(100)
var dummyOutputFee = uint64(100)
var dummyWorkTime = "+3GMT 9:00 - 21:00"
var dummyDelimiter = uint64(2)

func TestOperateOutput(t *testing.T) {
	dummyName := string(calc.Rand()[0:8])
	dummyUserAdress := calc.Rand()
	dummyMarketAdress := calc.Rand()
	user.Create(
		dummyUserAdress,
		dummyMessageKey,
		dummyName,
	)
	output := trade.Output{
		Adress: dummyUserAdress,
		Main:   100,
		Market: 50,
	}
	operateOutput(output, dummyMarketAdress)
	u := user.Get(dummyUserAdress)
	if u.Balance != 100 || u.Balances[string(dummyMarketAdress)] != 50 {
		t.Error("output was operated incorrectly")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(dummyUserAdress)
}

func TestAttachUnbounededBuys(t *testing.T) {
	adress := calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	Create(
		adress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	mkt := Get(adress)
	badBuy := trade.Buy{
		Offer:   100,
		Recieve: 100,
	}
	attached1 := mkt.AttachBuy(&badBuy)
	if attached1 {
		t.Error("trade without market attachment should not be attached")
	}
	lookedMkt := Look(adress)
	attached2 := lookedMkt.AttachBuy(&badBuy)
	if attached2 {
		t.Error("trade should not be attached to looked market")
	}
	data.TestRM(adress)
	data.TestRM([]byte(dummyName))
}

func TestAttachUnbounededSells(t *testing.T) {
	adress := calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	Create(
		adress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	mkt := Get(adress)
	badSell := trade.Sell{
		Offer:   100,
		Recieve: 100,
	}
	attached1 := mkt.AttachSell(&badSell)
	if attached1 {
		t.Error("trade without market attachment should not be attached")
	}
	lookedMkt := Look(adress)
	attached2 := lookedMkt.AttachSell(&badSell)
	if attached2 {
		t.Error("trade should not be attached to looked market")
	}
	data.TestRM(adress)
	data.TestRM([]byte(dummyName))
}

func TestAttachAndOperateOutputs(t *testing.T) {
	marketAdress := calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	Create(
		marketAdress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	mkt := Get(marketAdress)
	sellerAdress := calc.Rand()
	sellerName := string(calc.Rand()[0:8])
	user.Create(
		sellerAdress,
		dummyMessageKey,
		sellerName,
	)
	seller := user.Get(sellerAdress)
	seller.Balances[string(marketAdress)] = 100
	sellerAdress2 := calc.Rand()
	sellerName2 := string(calc.Rand()[0:8])
	user.Create(
		sellerAdress2,
		dummyMessageKey,
		sellerName2,
	)
	seller2 := user.Get(sellerAdress2)
	seller.Balances[string(marketAdress)] = 100
	buyerAdress := calc.Rand()
	buyerName := string(calc.Rand()[0:8])
	user.Create(
		buyerAdress,
		dummyMessageKey,
		buyerName,
	)
	buyer := user.Get(buyerAdress)
	buyer.Balance = 200
	sell := trade.Sell{
		Offer:   100,
		Recieve: 100,
	}
	buy := trade.Buy{
		Offer:   200,
		Recieve: 200,
	}
	sell2 := trade.Sell{
		Offer:   100,
		Recieve: 100,
	}
	buyer.AttachBuy(&buy)
	seller.AttachSell(&sell, marketAdress)
	seller2.AttachSell(&sell2, marketAdress)
	mkt.AttachSell(&sell)
	mkt.AttachBuy(&buy)
	mkt.AttachSell(&sell2)
	time.Sleep(time.Second)
}
