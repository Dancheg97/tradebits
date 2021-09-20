package market

import (
	"reflect"
	"sync_tree/calc"
	"sync_tree/data"
	"testing"
)

var dummyMessageKey = []byte{1, 2, 3, 4, 5}
var dummyImageLink = "test.imagelink/thereisnoimagebythislink"
var dummyName = "Test Market Name"
var dummyDescription = `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`
var dummyInputFee = uint64(100)
var dummyOutputFee = uint64(100)
var dummyWorkTime = "+3GMT 9:00 - 21:00"
var dummyDelimiter = uint64(2)

func TestCreateNewMarket(t *testing.T) {
	var dummyAdress = calc.Rand()
	err := Create(
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
	if err != nil {
		t.Error(err)
	}
	data.TestRM(dummyAdress)
}

func TestCreateMarketBadAdress(t *testing.T) {
	var badAdress = []byte{0, 1, 2, 3}
	err := Create(
		badAdress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if err == nil {
		t.Error("market with bad adress should not be created")
		data.TestRM(badAdress)
	}
}

func TestCreateMarketBadName(t *testing.T) {
	var dummyAdress = calc.Rand()
	err := Create(
		dummyAdress,
		"ola",
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if err == nil {
		t.Error("market with name that small should not be created")
		data.TestRM(dummyAdress)
	}
}

func TestCreateMarketBadDescription(t *testing.T) {
	var dummyAdress = calc.Rand()
	err := Create(
		dummyAdress,
		dummyName,
		dummyMessageKey,
		"coca cola",
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if err == nil {
		t.Error("market with description that small should not be created")
		data.TestRM(dummyAdress)
	}
}

func TestCreateMarketBadFee(t *testing.T) {
	var dummyAdress = calc.Rand()
	err := Create(
		dummyAdress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		502,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if err == nil {
		t.Error("market with fee that big should not be created")
		data.TestRM(dummyAdress)
	}
}

func TestCreateMarketBadWorkTime(t *testing.T) {
	var dummyAdress = calc.Rand()
	err := Create(
		dummyAdress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		"9-21",
		dummyDelimiter,
	)
	if err == nil {
		t.Error("market with work time that small should not be created")
		data.TestRM(dummyAdress)
	}
}

func TestCreateMarketBadDelimited(t *testing.T) {
	var dummyAdress = calc.Rand()
	err := Create(
		dummyAdress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		42,
	)
	if err == nil {
		t.Error("market with delimiter that big should not be created")
		data.TestRM(dummyAdress)
	}
}

func TestCreateExistingMarket(t *testing.T) {
	var adress = calc.Rand()
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
	err := Create(
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
	if err == nil {
		t.Error("new market should not be craeted")
	}
	data.TestRM(adress)
}

func TestGetFreeMarket(t *testing.T) {
	var adress = calc.Rand()
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
	market := Get(adress)
	if !reflect.DeepEqual(market.Name, dummyName) {
		t.Error("keys are not the same, get asset error")
	}
	market.Save()
	data.TestRM(adress)
}

func TestGetMarketThatDontExist(t *testing.T) {
	var adress = calc.Rand()
	mkt := Get(adress)
	if mkt != nil {
		t.Error("that test should not return anything")
	}
}

// func TestMarketLook(t *testing.T) {
// 	var adress = calc.Rand()
// 	Create(
// 		adress,
// 		dummyName,
// 		dummyMessageKey,
// 		dummyDescription,
// 		dummyImageLink,
// 		dummyInputFee,
// 		dummyOutputFee,
// 		dummyWorkTime,
// 		dummyDelimiter,
// 	)
// 	mkt := Look(adress)
// 	if !reflect.DeepEqual(mkt.MesKey, mesKey) {
// 		t.Error("keys are not the same, look asset error")
// 	}
// 	data.TestRM(adress)
// }

// func TestAttachUnboundedTrades(t *testing.T) {
// 	var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 229}
// 	var mesKey = []byte{1, 2, 3, 4, 5}
// 	Create(adress, name, mesKey, descr, img, inpFee, outFee, workTime)
// 	mkt := Get(adress)
// 	sell := trade.Sell{
// 		Offer:   100,
// 		Recieve: 100,
// 	}
// 	buy := trade.Buy{
// 		Offer:   100,
// 		Recieve: 100,
// 	}
// 	buyAttached := mkt.AttachBuy(&buy)
// 	sellAttached := mkt.AttachSell(&sell)
// 	if buyAttached || sellAttached {
// 		t.Error("those trades should not be attached cuz they are unbounded")
// 	}
// 	data.TestRM(adress)
// }

// func TestAttachToLookedMarket(t *testing.T) {
// 	var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 119, 120, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 229}
// 	var mesKey = []byte{1, 2, 3, 4, 5}
// 	Create(adress, name, mesKey, descr, img, inpFee, outFee, workTime)
// 	mkt := Look(adress)
// 	sell := trade.Sell{
// 		Offer:   100,
// 		Recieve: 100,
// 	}
// 	buy := trade.Buy{
// 		Offer:   100,
// 		Recieve: 100,
// 	}
// 	buyAttached := mkt.AttachBuy(&buy)
// 	sellAttached := mkt.AttachSell(&sell)
// 	if buyAttached || sellAttached {
// 		t.Error("those trades should not be attached cuz they are unbounded")
// 	}
// 	data.TestRM(adress)
// }

// func TestAttachSingleNormalBuy(t *testing.T) {
// 	var marketAdress = []byte{1, 2, 33, 42, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 119, 120, 21, 22, 23, 124, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 122, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 161, 62, 63, 229}
// 	var marketMesKey = []byte{1, 2, 3, 4, 5}
// 	var marketImg = "asset image link . example"
// 	var name = "newAsset"
// 	var descr = "descrx"
// 	Create(marketAdress, name, marketMesKey, descr, marketImg, inpFee, outFee, workTime)
// 	mkt := Get(marketAdress)
// 	var adress = []byte{1, 22, 3, 44, 5, 16, 7, 8, 9, 10, 11, 112, 13, 14, 15, 16, 19, 18, 19, 20, 21, 122, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 121, 59, 22, 91, 91, 91, 91}
// 	var mesKey = []byte{1, 2, 3, 4, 5}
// 	var img = "user image link"
// 	user.Create(adress, mesKey, img)
// 	usr := user.Get(adress)
// 	usr.Balance = 100
// 	buy := trade.Buy{
// 		Offer:   100,
// 		Recieve: 100,
// 	}
// 	attachedToUser := usr.AttachBuy(&buy)
// 	if !attachedToUser {
// 		t.Error("trade should be attached to user")
// 	}
// 	attachedToMarket := mkt.AttachBuy(&buy)
// 	if !attachedToMarket {
// 		t.Error("trade should be attached to market")
// 	}
// 	data.TestRM(marketAdress)
// 	data.TestRM(adress)
// }

// func TestAttachSingleNormalSell(t *testing.T) {
// 	var marketAdress = []byte{11, 2, 33, 42, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 119, 120, 121, 22, 23, 124, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 122, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 161, 62, 63, 229}
// 	var marketMesKey = []byte{1, 2, 3, 4, 5}
// 	var marketImg = "asset image link . example"
// 	var name = "newAsset"
// 	var descr = "descrx"
// 	Create(marketAdress, name, marketMesKey, descr, marketImg, inpFee, outFee, workTime)
// 	mkt := Get(marketAdress)
// 	var adress = []byte{1, 22, 13, 44, 5, 16, 7, 8, 9, 10, 11, 112, 13, 14, 15, 16, 19, 18, 19, 20, 21, 122, 123, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 40, 21, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 121, 59, 22, 91, 91, 91, 91}
// 	var mesKey = []byte{1, 2, 3, 4, 5}
// 	var img = "user image link"
// 	user.Create(adress, mesKey, img)
// 	usr := user.Get(adress)
// 	usr.Balances[string(marketAdress)] = 100
// 	buy := trade.Sell{
// 		Offer:   100,
// 		Recieve: 100,
// 	}
// 	attachedToUser := usr.AttachSell(&buy, marketAdress)
// 	if !attachedToUser {
// 		t.Error("trade should be attached to user")
// 	}
// 	attachedToMarket := mkt.AttachSell(&buy)
// 	if !attachedToMarket {
// 		t.Error("trade should be attached to market")
// 	}
// 	data.TestRM(marketAdress)
// 	data.TestRM(adress)
// }

// func TestTwoUserTradesWithSameOffers(t *testing.T) {
// 	var marketAdress = []byte{1, 2, 33, 42, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 119, 1, 121, 22, 23, 124, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 111, 38, 1, 40, 41, 122, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 161, 62, 63, 229}
// 	var marketMesKey = []byte{1, 2, 3, 4, 5}
// 	var marketImg = "asset image link . example"
// 	var name = "newAsset"
// 	var descr = "descrx"
// 	Create(marketAdress, name, marketMesKey, descr, marketImg, inpFee, outFee, workTime)
// 	mkt := Get(marketAdress)

// 	var adress1 = []byte{121, 22, 13, 44, 5, 16, 7, 8, 9, 10, 11, 112, 13, 14, 15, 16, 19, 18, 19, 20, 121, 122, 123, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 40, 21, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 121, 59, 22, 91, 91, 91, 91}
// 	var mesKey1 = []byte{1, 2, 3, 4, 5}
// 	var img1 = "user image link"
// 	user.Create(adress1, mesKey1, img1)
// 	usr1 := user.Get(adress1)
// 	usr1.Balance = 100
// 	buy := trade.Buy{
// 		Offer:   100,
// 		Recieve: 100,
// 	}
// 	usr1.AttachBuy(&buy)
// 	mkt.AttachBuy(&buy)
// 	usr1.Save()

// 	var adress2 = []byte{129, 22, 13, 44, 5, 16, 7, 8, 9, 10, 110, 112, 13, 14, 15, 16, 19, 18, 19, 20, 21, 122, 123, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 140, 21, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 157, 121, 59, 22, 91, 91, 91, 91}
// 	var mesKey2 = []byte{1, 2, 3, 4, 5}
// 	var img2 = "user image link"
// 	user.Create(adress2, mesKey2, img2)
// 	usr2 := user.Get(adress2)
// 	usr2.Balances[string(marketAdress)] = 100
// 	sell := trade.Sell{
// 		Offer:   100,
// 		Recieve: 100,
// 	}
// 	usr2.AttachSell(&sell, marketAdress)
// 	mkt.AttachSell(&sell)
// 	usr2.Save()

// 	time.Sleep(time.Second * 1)

// 	usr1check := user.Look(adress1)
// 	if usr1check.Balance != 0 {
// 		t.Error("first user main balance fshould be equal to zero")
// 	}
// 	if usr1check.Balances[string(marketAdress)] != 100 {
// 		t.Error("first user market balance should be equal to 100")
// 	}
// 	usr2check := user.Look(adress2)
// 	if usr2check.Balances[string(marketAdress)] != 0 {
// 		t.Error("market balance of second user should be equal to zero")
// 	}
// 	if usr2check.Balance != 100 {
// 		t.Error("second user main balance should be equal to 100")
// 	}

// 	data.TestRM(adress1)
// 	data.TestRM(adress2)
// 	data.TestRM(marketAdress)
// }

// func TestFourUserTradesWithRandomOffers(t *testing.T) {
// 	var marketAdress = []byte{129, 22, 13, 44, 5, 16, 7, 8, 9, 10, 110, 112, 13, 14, 15, 16, 19, 18, 19, 20, 21, 122, 123, 1, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 140, 21, 1, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 157, 121, 59, 1, 91, 91, 91, 91}
// 	var mesKey = []byte{1, 2, 3, 4, 5}
// 	var img = "asset image link . example"
// 	Create(marketAdress, img, mesKey, img, img, inpFee, outFee, workTime)
// 	mkt := Get(marketAdress)

// 	var firstUserAdress = []byte{129, 22, 13, 44, 5, 16, 7, 8, 9, 10, 110, 112, 13, 14, 15, 16, 19, 18, 19, 20, 21, 122, 123, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 140, 21, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 157, 121, 59, 22, 1, 1, 91, 91}
// 	user.Create(firstUserAdress, mesKey, img)
// 	firstUser := user.Get(firstUserAdress)
// 	firstUser.Balance = 300
// 	firstUserTrade := trade.Buy{
// 		Offer:   270,
// 		Recieve: 130,
// 	}
// 	firstUser.AttachBuy(&firstUserTrade)
// 	mkt.AttachBuy(&firstUserTrade)
// 	firstUser.Save()

// 	var secondUserAdress = []byte{129, 22, 13, 44, 5, 16, 7, 8, 9, 10, 110, 112, 13, 14, 15, 16, 19, 18, 19, 20, 21, 1, 123, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 1, 21, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 157, 121, 1, 22, 91, 91, 91, 91}
// 	user.Create(secondUserAdress, mesKey, img)
// 	secondUser := user.Get(secondUserAdress)
// 	secondUser.Balances[string(marketAdress)] = 150
// 	secondUserTrade := trade.Sell{
// 		Offer:   80,
// 		Recieve: 130,
// 	}
// 	secondUser.AttachSell(&secondUserTrade, marketAdress)
// 	mkt.AttachSell(&secondUserTrade)
// 	secondUser.Save()

// 	var thirdUserAdress = []byte{1, 1, 1, 44, 5, 16, 7, 8, 9, 10, 110, 112, 13, 14, 15, 16, 19, 18, 19, 20, 21, 1, 123, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 1, 21, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 157, 121, 1, 22, 91, 91, 91, 91}
// 	user.Create(thirdUserAdress, mesKey, img)
// 	thirdUser := user.Get(thirdUserAdress)
// 	thirdUser.Balances[string(marketAdress)] = 150
// 	thirdUserTrade := trade.Sell{
// 		Offer:   20,
// 		Recieve: 15,
// 	}
// 	thirdUser.AttachSell(&thirdUserTrade, marketAdress)
// 	mkt.AttachSell(&thirdUserTrade)
// 	thirdUser.Save()

// 	time.Sleep(time.Second * 1)
// 	firstUserCheck := user.Look(firstUserAdress)
// 	if firstUserCheck.Balance != 30 {
// 		t.Error("first user balance should be equal to 30")
// 	}
// 	if firstUserCheck.Balances[string(marketAdress)] != 100 {
// 		t.Error("first user market balance should be equal to 100")
// 	}
// 	secondUserCheck := user.Look(secondUserAdress)
// 	if secondUserCheck.Balance != 130 {
// 		t.Error("second user balance should be equal to 130")
// 	}
// 	if secondUserCheck.Balances[string(marketAdress)] != 70 {
// 		t.Error("second user market balance should be equal to 70")
// 	}
// 	if len(mkt.Pool.Buys) != 1 {
// 		t.Error("market pool length should be equal to one")
// 	}
// 	if mkt.Pool.Buys[0].Offer != 125 {
// 		t.Error("current offer of market buy should be equal to 125")
// 		t.Error(mkt.Pool.Buys[0].Offer)
// 	}
// 	if mkt.Pool.Buys[0].Recieve != 30 {
// 		t.Error("current offer of market buy should be equal to 30")
// 		t.Error(mkt.Pool.Buys[0].Recieve)
// 	}
// 	if len(mkt.Pool.Sells) != 0 {
// 		t.Error("there should not be any active market sell")
// 	}
// 	if len(mkt.Pool.Outputs) != 0 {
// 		t.Error("there should not be any market outputs")
// 	}
// 	data.TestRM(marketAdress)
// 	data.TestRM(firstUserAdress)
// 	data.TestRM(secondUserAdress)
// }

// func TestAttachFirstlySellThanBuy(t *testing.T) {
// 	var marketAdress = []byte{1, 22, 13, 44, 5, 16, 7, 8, 9, 10, 110, 112, 13, 14, 15, 16, 19, 18, 19, 20, 1, 122, 123, 1, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 1, 39, 140, 21, 1, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 1, 1, 1, 59, 1, 91, 91, 91, 91}
// 	var mesKey = []byte{1, 2, 3, 4, 5}
// 	var img = "asset image link . example"
// 	Create(marketAdress, img, mesKey, img, img, inpFee, outFee, workTime)
// 	mkt := Get(marketAdress)

// 	var firstUserAdress = []byte{1, 22, 13, 44, 5, 16, 7, 8, 9, 10, 110, 112, 13, 14, 15, 16, 19, 18, 19, 1, 21, 122, 123, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 1, 39, 140, 21, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 2, 2, 121, 59, 22, 1, 1, 91, 91}
// 	user.Create(firstUserAdress, mesKey, img)
// 	firstUser := user.Get(firstUserAdress)
// 	firstUser.Balance = 300
// 	firstUserTrade := trade.Buy{
// 		Offer:   270,
// 		Recieve: 130,
// 	}
// 	firstUser.AttachBuy(&firstUserTrade)
// 	firstUser.Save()

// 	var secondUserAdress = []byte{1, 22, 13, 44, 5, 16, 7, 8, 9, 10, 110, 2, 13, 14, 15, 16, 19, 18, 19, 20, 2, 1, 123, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 1, 1, 21, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 1, 2, 1, 22, 91, 91, 91, 91}
// 	user.Create(secondUserAdress, mesKey, img)
// 	secondUser := user.Get(secondUserAdress)
// 	secondUser.Balances[string(marketAdress)] = 150
// 	secondUserTrade := trade.Sell{
// 		Offer:   80,
// 		Recieve: 130,
// 	}
// 	secondUser.AttachSell(&secondUserTrade, marketAdress)
// 	secondUser.Save()

// 	mkt.AttachSell(&secondUserTrade)
// 	mkt.AttachBuy(&firstUserTrade)

// 	time.Sleep(time.Second)

// 	firstUserCheck := user.Look(firstUserAdress)
// 	secondUserCheck := user.Look(secondUserAdress)
// 	if firstUserCheck.Balance != 30 {
// 		t.Error("first user balance should be equal to 30")
// 	}
// 	if firstUserCheck.Balances[string(marketAdress)] != 80 {
// 		t.Error("first user market balance should be equal to 80")
// 	}
// 	if secondUserCheck.Balance != 130 {
// 		t.Error("second user main balance should be equal to 130")
// 	}
// 	if secondUserCheck.Balances[string(marketAdress)] != 70 {
// 		t.Error("second user market balance should be equal to 70")
// 	}
// 	if mkt.Pool.Buys[0].Offer != 140 {
// 		t.Error("active buy offer should be equal tp 140")
// 	}
// 	if mkt.Pool.Buys[0].Recieve != 50 {
// 		t.Error("acrive buy recieve should be equal to 50")
// 	}

// 	data.TestRM(marketAdress)
// 	data.TestRM(firstUserAdress)
// 	data.TestRM(secondUserAdress)
// }

// func TestIfUserHasTrdadesWhenHeHaveSome(t *testing.T) {
// 	var marketAdress = []byte{1, 22, 13, 44, 5, 16, 7, 8, 9, 10, 110, 112, 13, 14, 15, 16, 19, 18, 19, 20, 1, 3, 4, 1, 5, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 1, 39, 2, 21, 1, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 1, 1, 1, 1, 1, 91, 91, 91, 91}
// 	var mesKey = []byte{1, 2, 3, 4, 5}
// 	var img = "asset image link . example"
// 	Create(marketAdress, img, mesKey, img, img, inpFee, outFee, workTime)
// 	mkt := Get(marketAdress)

// 	sell := trade.Sell{
// 		Adress:  []byte{0},
// 		Offer:   1,
// 		Recieve: 5,
// 	}
// 	mkt.AttachSell(&sell)

// 	userHaveSell := mkt.HasTrades([]byte{0})
// 	if !userHaveSell {
// 		t.Error("error, this user should have some sell trade")
// 	}

// 	buy := trade.Buy{
// 		Adress:  []byte{0},
// 		Offer:   1,
// 		Recieve: 5,
// 	}
// 	mkt.AttachBuy(&buy)
// 	userHaveBuy := mkt.HasTrades([]byte{0})
// 	if !userHaveBuy {
// 		t.Error("error, this user should have some buy trade")
// 	}

// 	data.TestRM(marketAdress)
// }

// func TestIfUserHasTradesWhenHeDont(t *testing.T) {
// 	var marketAdress = []byte{1, 22, 13, 44, 5, 16, 7, 8, 9, 10, 2, 112, 13, 14, 15, 16, 19, 18, 1, 1, 1, 3, 4, 1, 5, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 1, 1, 2, 1, 1, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 1, 1, 1, 1, 1, 91, 91, 2, 91}
// 	var mesKey = []byte{1, 2, 3, 4, 5}
// 	var img = "asset image link . example"
// 	Create(marketAdress, img, mesKey, img, img, inpFee, outFee, workTime)
// 	mkt := Get(marketAdress)

// 	userHaveSomeTrades := mkt.HasTrades([]byte{1})
// 	if userHaveSomeTrades {
// 		t.Error("there should not be any active trades for that adress")
// 	}

// 	data.TestRM(marketAdress)
// }

// func TestUserCancelBuy(t *testing.T) {
// 	var marketAdress = []byte{129, 22, 13, 44, 5, 16, 7, 8, 9, 10, 110, 112, 13, 14, 15, 16, 19, 18, 19, 1, 21, 122, 123, 1, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 1, 2, 21, 1, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 157, 121, 59, 1, 91, 91, 91, 91}
// 	var mesKey = []byte{1, 2, 3, 4, 5}
// 	var img = "asset image link . example"
// 	Create(marketAdress, img, mesKey, img, img, inpFee, outFee, workTime)
// 	mkt := Get(marketAdress)

// 	var userAdress = []byte{1, 22, 13, 44, 5, 16, 7, 8, 9, 10, 7, 112, 13, 14, 15, 16, 19, 18, 7, 20, 21, 1, 123, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 7, 39, 140, 21, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 7, 121, 59, 22, 1, 1, 91, 91}
// 	user.Create(userAdress, mesKey, img)
// 	usr := user.Get(userAdress)
// 	usr.Balance = 300
// 	usrBuy := trade.Buy{
// 		Offer:   270,
// 		Recieve: 130,
// 	}
// 	usr.AttachBuy(&usrBuy)
// 	mkt.AttachBuy(&usrBuy)
// 	usr.Save()

// 	mkt.CancelTrades(userAdress)

// 	time.Sleep(time.Second)
// 	userCheck := user.Look(userAdress)
// 	if len(mkt.Pool.Buys) != 0 {
// 		t.Error("the trade have not been cancelled, there should not be active trades on the market")
// 	}
// 	if userCheck.Balance != 300 {
// 		t.Error("user main balance should be equal to 300, cuz his trade has benn cancelled")
// 	}
// }

// func TestUserCancelSell(t *testing.T) {
// 	var marketAdress = []byte{1, 22, 13, 44, 5, 16, 7, 8, 9, 10, 110, 112, 13, 14, 15, 16, 19, 18, 19, 20, 1, 5, 2, 1, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 1, 39, 140, 2, 1, 43, 5, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 1, 1, 1, 59, 1, 5, 91, 91, 91}
// 	var mesKey = []byte{1, 2, 3, 4, 5}
// 	var img = "asset image link . example"
// 	Create(marketAdress, img, mesKey, img, img, inpFee, outFee, workTime)
// 	mkt := Get(marketAdress)

// 	var userAdress = []byte{1, 22, 13, 44, 5, 16, 7, 8, 9, 10, 110, 2, 13, 14, 15, 16, 19, 18, 19, 20, 2, 1, 1, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 1, 1, 21, 42, 43, 4, 5, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 1, 2, 1, 22, 91, 2, 91, 91}
// 	user.Create(userAdress, mesKey, img)
// 	usr := user.Get(userAdress)
// 	usr.Balances[string(marketAdress)] = 150
// 	sell := trade.Sell{
// 		Offer:   80,
// 		Recieve: 130,
// 	}
// 	usr.AttachSell(&sell, marketAdress)
// 	usr.Save()

// 	mkt.AttachSell(&sell)

// 	mkt.CancelTrades(userAdress)

// 	time.Sleep(time.Second)
// 	userCheck := user.Look(userAdress)
// 	if len(mkt.Pool.Sells) != 0 {
// 		t.Error("there should not be any active sell, because user cancelled that")
// 	}
// 	if userCheck.Balances[string(marketAdress)] != 150 {
// 		t.Error("user market balance should be equal to 150 after sell trade cancellation")
// 	}
// }
