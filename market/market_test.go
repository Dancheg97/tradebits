package market

import (
	"reflect"
	"sync_tree/data"
	"sync_tree/trade"
	"testing"
	"time"
)

func TestCreateNewMarket(t *testing.T) {
	var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 63}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "asset image link . example"
	var name = "newAsset"
	var descr = "descrx"
	err := Create(adress, name, mesKey, descr, img)
	if err != nil {
		t.Error("failed to craete new market")
	}
	data.TestRM(adress)
}

func TestCreateExistingMarket(t *testing.T) {
	var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 61}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "asset image link . example"
	var name = "newAsset"
	var descr = "descrx"
	Create(adress, name, mesKey, descr, img)
	err := Create(adress, name, mesKey, descr, img)
	if err == nil {
		t.Error("new market should not be craeted")
	}
	data.TestRM(adress)
}

func TestGetFreeMarket(t *testing.T) {
	var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 89}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "asset image link . example"
	var name = "newAsset"
	var descr = "descrx"
	Create(adress, name, mesKey, descr, img)
	market := Get(adress)
	defer market.Save()
	if !reflect.DeepEqual(market.MesKey, mesKey) {
		t.Error("keys are not the same, get asset error")
	}
	data.TestRM(adress)
}

func TestAssetLook(t *testing.T) {
	var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 12}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "asset image link . example"
	var name = "newAsset"
	var descr = "descrx"
	Create(adress, name, mesKey, descr, img)
	mkt := Look(adress)
	if !reflect.DeepEqual(mkt.MesKey, mesKey) {
		t.Error("keys are not the same, look asset error")
	}
	data.TestRM(adress)
}

var mkt2 *market

func getBusyMarket(adress []byte) {
	mkt2 = Get(adress)
}

func TestMarketGetAfterBusy(t *testing.T) {
	var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 129}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "asset image link . example"
	var name = "newAsset"
	var descr = "descrx"
	Create(adress, name, mesKey, descr, img)
	mkt1 := Get(adress)
	go getBusyMarket(adress)
	time.Sleep(time.Second)
	mkt1.Save()
	time.Sleep(time.Second)
	if !reflect.DeepEqual(mkt2.adress, adress) {
		t.Error("adresses should be equal")
	}
	mkt2.Save()
	data.TestRM(adress)
}

func TestMarketMessages(t *testing.T) {
	var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 229}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "asset image link . example"
	var name = "newAsset"
	var descr = "descrx"
	Create(adress, name, mesKey, descr, img)
	mkt := Get(adress)
	mkt.PutMessage([]byte{0, 1, 2}, "yo")
	msgs := mkt.GetAllMessages()
	if !reflect.DeepEqual(msgs[string([]byte{0, 1, 2})], "yo") {
		t.Error("message should be 'yo'")
	}
	mkt.Save()
	data.TestRM(adress)
}

func TestAttachUnboundedTrades(t *testing.T) {
	var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 229}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "asset image link . example"
	var name = "newAsset"
	var descr = "descrx"
	Create(adress, name, mesKey, descr, img)
	mkt := Get(adress)
	sell := trade.Sell{
		Offer:   100,
		Recieve: 100,
	}
	buy := trade.Buy{
		Offer:   100,
		Recieve: 100,
	}
	buyAttached := mkt.AttachBuy(buy)
	sellAttached := mkt.AttachSell(sell)
	if buyAttached || sellAttached {
		t.Error("those trades should not be attached cuz they are unbounded")
	}
	data.TestRM(adress)
}

func TestAttachToLookedMarket(t *testing.T) {
	var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 119, 120, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 229}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "asset image link . example"
	var name = "newAsset"
	var descr = "descrx"
	Create(adress, name, mesKey, descr, img)
	mkt := Look(adress)
	sell := trade.Sell{
		Offer:   100,
		Recieve: 100,
	}
	buy := trade.Buy{
		Offer:   100,
		Recieve: 100,
	}
	buyAttached := mkt.AttachBuy(buy)
	sellAttached := mkt.AttachSell(sell)
	if buyAttached || sellAttached {
		t.Error("those trades should not be attached cuz they are unbounded")
	}
	data.TestRM(adress)
}
