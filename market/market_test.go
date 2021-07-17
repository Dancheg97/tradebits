package market

import (
	"reflect"
	"sync_tree/data"
	"sync_tree/trade"
	"sync_tree/user"
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
	buyAttached := mkt.AttachBuy(&buy)
	sellAttached := mkt.AttachSell(&sell)
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
	buyAttached := mkt.AttachBuy(&buy)
	sellAttached := mkt.AttachSell(&sell)
	if buyAttached || sellAttached {
		t.Error("those trades should not be attached cuz they are unbounded")
	}
	data.TestRM(adress)
}

func TestAttachSingleNormalBuy(t *testing.T) {
	var marketAdress = []byte{1, 2, 33, 42, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 119, 120, 21, 22, 23, 124, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 122, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 161, 62, 63, 229}
	var marketMesKey = []byte{1, 2, 3, 4, 5}
	var marketImg = "asset image link . example"
	var name = "newAsset"
	var descr = "descrx"
	Create(marketAdress, name, marketMesKey, descr, marketImg)
	mkt := Get(marketAdress)
	var adress = []byte{1, 22, 3, 44, 5, 16, 7, 8, 9, 10, 11, 112, 13, 14, 15, 16, 19, 18, 19, 20, 21, 122, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 121, 59, 22, 91, 91, 91, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	user.Create(adress, mesKey, img)
	usr := user.Get(adress)
	usr.Balance = 100
	buy := trade.Buy{
		Offer:   100,
		Recieve: 100,
	}
	attachedToUser := usr.AttachBuy(&buy)
	if !attachedToUser {
		t.Error("trade should be attached to user")
	}
	attachedToMarket := mkt.AttachBuy(&buy)
	if !attachedToMarket {
		t.Error("trade should be attached to market")
	}
	data.TestRM(marketAdress)
	data.TestRM(adress)
}

func TestAttachSingleNormalSell(t *testing.T) {
	var marketAdress = []byte{11, 2, 33, 42, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 119, 120, 121, 22, 23, 124, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 122, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 161, 62, 63, 229}
	var marketMesKey = []byte{1, 2, 3, 4, 5}
	var marketImg = "asset image link . example"
	var name = "newAsset"
	var descr = "descrx"
	Create(marketAdress, name, marketMesKey, descr, marketImg)
	mkt := Get(marketAdress)
	var adress = []byte{1, 22, 13, 44, 5, 16, 7, 8, 9, 10, 11, 112, 13, 14, 15, 16, 19, 18, 19, 20, 21, 122, 123, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 40, 21, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 121, 59, 22, 91, 91, 91, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	user.Create(adress, mesKey, img)
	usr := user.Get(adress)
	usr.Markets[string(marketAdress)] = 100
	buy := trade.Sell{
		Offer:   100,
		Recieve: 100,
	}
	attachedToUser := usr.AttachSell(&buy, marketAdress)
	if !attachedToUser {
		t.Error("trade should be attached to user")
	}
	attachedToMarket := mkt.AttachSell(&buy)
	if !attachedToMarket {
		t.Error("trade should be attached to market")
	}
	data.TestRM(marketAdress)
	data.TestRM(adress)
}

func TestTwoUserTradesWithSameOffers(t *testing.T) {
	var marketAdress = []byte{1, 2, 33, 42, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 119, 1, 121, 22, 23, 124, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 111, 38, 1, 40, 41, 122, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 161, 62, 63, 229}
	var marketMesKey = []byte{1, 2, 3, 4, 5}
	var marketImg = "asset image link . example"
	var name = "newAsset"
	var descr = "descrx"
	Create(marketAdress, name, marketMesKey, descr, marketImg)
	mkt := Get(marketAdress)

	var adress1 = []byte{121, 22, 13, 44, 5, 16, 7, 8, 9, 10, 11, 112, 13, 14, 15, 16, 19, 18, 19, 20, 121, 122, 123, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 40, 21, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 121, 59, 22, 91, 91, 91, 91}
	var mesKey1 = []byte{1, 2, 3, 4, 5}
	var img1 = "user image link"
	user.Create(adress1, mesKey1, img1)
	usr1 := user.Get(adress1)
	usr1.Balance = 100
	buy := trade.Buy{
		Offer:   100,
		Recieve: 100,
	}
	usr1.AttachBuy(&buy)
	mkt.AttachBuy(&buy)
	usr1.Save()

	var adress2 = []byte{129, 22, 13, 44, 5, 16, 7, 8, 9, 10, 110, 112, 13, 14, 15, 16, 19, 18, 19, 20, 21, 122, 123, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 140, 21, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 157, 121, 59, 22, 91, 91, 91, 91}
	var mesKey2 = []byte{1, 2, 3, 4, 5}
	var img2 = "user image link"
	user.Create(adress2, mesKey2, img2)
	usr2 := user.Get(adress2)
	usr2.Markets[string(marketAdress)] = 100
	sell := trade.Sell{
		Offer:   100,
		Recieve: 100,
	}
	usr2.AttachSell(&sell, marketAdress)
	mkt.AttachSell(&sell)
	usr2.Save()

	time.Sleep(time.Second * 1)

	usr1check := user.Look(adress1)
	if usr1check.Balance != 0 {
		t.Error("first user main balance fshould be equal to zero")
	}
	if usr1check.Markets[string(marketAdress)] != 100 {
		t.Error("first user market balance should be equal to 100")
	}
	usr2check := user.Look(adress2)
	if usr2check.Markets[string(marketAdress)] != 0 {
		t.Error("market balance of second user should be equal to zero")
	}
	if usr2check.Balance != 100 {
		t.Error("second user main balance should be equal to 100")
	}

	data.TestRM(adress1)
	data.TestRM(adress2)
	data.TestRM(marketAdress)
}

func TestFourUserTradesWithRandomOffers(t *testing.T) {
	var marketAdress = []byte{129, 22, 13, 44, 5, 16, 7, 8, 9, 10, 110, 112, 13, 14, 15, 16, 19, 18, 19, 20, 21, 122, 123, 1, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 140, 21, 1, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 157, 121, 59, 1, 91, 91, 91, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "asset image link . example"
	Create(marketAdress, img, mesKey, img, img)
	mkt := Get(marketAdress)

	var firstUserAdress = []byte{129, 22, 13, 44, 5, 16, 7, 8, 9, 10, 110, 112, 13, 14, 15, 16, 19, 18, 19, 20, 21, 122, 123, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 140, 21, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 157, 121, 59, 22, 1, 1, 91, 91}
	user.Create(firstUserAdress, mesKey, img)
	firstUser := user.Get(firstUserAdress)
	firstUser.Balance = 300
	firstUserTrade := trade.Buy{
		Offer:   270,
		Recieve: 130,
	}
	firstUser.AttachBuy(&firstUserTrade)
	mkt.AttachBuy(&firstUserTrade)
	firstUser.Save()

	var secondUserAdress = []byte{129, 22, 13, 44, 5, 16, 7, 8, 9, 10, 110, 112, 13, 14, 15, 16, 19, 18, 19, 20, 21, 1, 123, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 1, 21, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 157, 121, 1, 22, 91, 91, 91, 91}
	user.Create(secondUserAdress, mesKey, img)
	secondUser := user.Get(secondUserAdress)
	secondUser.Markets[string(marketAdress)] = 150
	secondUserTrade := trade.Sell{
		Offer:   80,
		Recieve: 130,
	}
	secondUser.AttachSell(&secondUserTrade, marketAdress)
	mkt.AttachSell(&secondUserTrade)
	secondUser.Save()

	time.Sleep(time.Second * 1)
	firstUserCheck := user.Look(firstUserAdress)
	t.Error(firstUserCheck.Balance)
	t.Error(firstUserCheck.Markets[string(marketAdress)])
	secondUserCheck := user.Look(secondUserAdress)
	t.Error(secondUserCheck.Balance)
	t.Error(secondUserCheck.Markets[string(marketAdress)])
	t.Error(mkt.Pool.Buys)
	t.Error(mkt.Pool.Sells)
	t.Error(mkt.Pool.Outputs)

	data.TestRM(marketAdress)
	data.TestRM(firstUserAdress)
	data.TestRM(secondUserAdress)
}
