package user

import (
	"reflect"
	"sync_tree/data"
	"sync_tree/trade"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 91, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	err := Create(adress, mesKey, img)
	if err != nil {
		t.Error("attemt to create new user failed")
	}
	data.TestRM(adress)
}

func TestCreateExisting(t *testing.T) {
	var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 91, 91, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	err := Create(adress, mesKey, img)
	if err == nil {
		t.Error("attemt to create existing user succeded, that is bad error")
	}
	data.TestRM(adress)
}

func TestGetFreeUser(t *testing.T) {
	var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 22, 91, 91, 91, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	freeUser := Get(adress)
	freeUser.Save()
	if freeUser.PublicName != "user image link" {
		t.Error("get free user error")
	}
	data.TestRM(adress)
}

var usr2 *user

func getBusyUser(adress []byte) {
	usr2 = Get(adress)
}

func TestGetBusyUser(t *testing.T) {
	var adress = []byte{1, 22, 3, 44, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 22, 91, 91, 91, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr1 := Get(adress)
	go getBusyUser(adress)
	time.Sleep(time.Second)
	usr1.Save()
	time.Sleep(time.Second)
	if !reflect.DeepEqual(usr2.adress, adress) {
		t.Error("adress of second user should be the same")
	}
	data.TestRM(adress)
}

func TestUserLook(t *testing.T) {
	var adress = []byte{1, 22, 3, 44, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 121, 59, 22, 91, 91, 91, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr := Look(adress)
	if len(usr.adress) != 0 {
		t.Error("user adress should be empty")
	}
	if usr.PublicName != img {
		t.Error("user info is incorrect")
	}
	data.TestRM(adress)
}

func TestUserMessages(t *testing.T) {
	var adress = []byte{1, 22, 3, 44, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 121, 59, 22, 91, 91, 91, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr := Get(adress)
	usr.PutMessage([]byte{0, 1, 2}, "yo")
	usr.Save()
	usrTake := Look(adress)
	msgs := usrTake.GetAllMessages()
	if len(msgs) != 1 {
		t.Error("there should be only one message")
	}
	if msgs[string([]byte{0, 1, 2})] == "yo" {
		t.Error("values should be the same")
	}
	data.TestRM(adress)
}

func TestAttachBuyToLookedUser(t *testing.T) {
	var adress = []byte{1, 22, 3, 44, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 59, 56, 99, 121, 59, 22, 91, 91, 91, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr := Look(adress)
	buy := trade.Buy{}
	err := usr.AttachBuy(buy)
	if err == nil {
		t.Error("this test should throw an error, because it should be impossible to attach trade to unsavable user")
	}
	data.TestRM(adress)
}

func TestAttachBuyWithZeroOffer(t *testing.T) {
	var adress = []byte{1, 22, 3, 44, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 19, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 121, 59, 22, 91, 91, 91, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr := Get(adress)
	buy := trade.Buy{
		Offer:   0,
		Recieve: 1000,
	}
	err := usr.AttachBuy(buy)
	if err == nil {
		t.Error("this test should throw an error, cuz trade is 0 offer")
	}
	data.TestRM(adress)
}

func TestAttachToBigBuyTrade(t *testing.T) {
	var adress = []byte{1, 22, 3, 44, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 19, 18, 19, 20, 21, 22, 23, 24, 25, 232, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 121, 59, 22, 91, 91, 232, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr := Get(adress)
	buy := trade.Buy{
		Offer:   1000,
		Recieve: 1000,
	}
	err := usr.AttachBuy(buy)
	if err == nil {
		t.Error("this test should throw an error, cuz user has no money")
	}
	data.TestRM(adress)
}

func TestAttachNormalBuy(t *testing.T) {
	var adress = []byte{1, 22, 3, 44, 15, 6, 7, 8, 9, 110, 11, 12, 13, 14, 15, 16, 19, 18, 19, 20, 21, 22, 23, 24, 25, 232, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 87, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 121, 59, 22, 91, 91, 232, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr := Get(adress)
	usr.Balance = 1000
	buy := trade.Buy{
		Offer:   1000,
		Recieve: 1000,
	}
	err := usr.AttachBuy(buy)
	if err != nil {
		t.Error("this test should not throw any errors cuz its fine with user")
	}
	data.TestRM(adress)
}

func TestAttachSellToLookedUser(t *testing.T) {
	var adress = []byte{1, 22, 32, 244, 5, 6, 7, 18, 9, 110, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 136, 37, 38, 139, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 59, 56, 99, 121, 59, 22, 91, 91, 91, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr := Look(adress)
	sell := trade.Sell{}
	err := usr.AttachSell(sell)
	if err == nil {
		t.Error("this test should throw an error, because it should be impossible to attach trade to unsavable user")
	}
	data.TestRM(adress)
}

func TestAttachSellWithZeroOffer(t *testing.T) {
	var adress = []byte{1, 22, 13, 44, 15, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 19, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 111, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 155, 156, 157, 121, 59, 22, 91, 91, 91, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr := Get(adress)
	sell := trade.Sell{
		Offer:   0,
		Recieve: 1000,
	}
	err := usr.AttachSell(sell)
	if err == nil {
		t.Error("this test should throw an error, cuz trade is 0 offer")
	}
	data.TestRM(adress)
}

func TestAttachToBigSellTrade(t *testing.T) {
	var adress = []byte{1, 22, 3, 44, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 19, 18, 19, 20, 21, 22, 23, 24, 25, 232, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 40, 141, 142, 143, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 121, 59, 22, 91, 91, 232, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr := Get(adress)
	usr.Markets["0"] = 0
	sell := trade.Sell{
		Offer:   1000,
		Recieve: 1000,
		Adress:  []byte("0"),
	}
	err := usr.AttachSell(sell)
	if err == nil {
		t.Error("this test should throw an error, cuz user has no money")
	}
	data.TestRM(adress)
}

func TestAttachNormalSell(t *testing.T) {
	var adress = []byte{1, 22, 3, 44, 5, 16, 7, 8, 9, 110, 11, 12, 13, 14, 15, 16, 19, 18, 19, 20, 21, 22, 23, 24, 25, 232, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 87, 40, 41, 42, 143, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 121, 59, 22, 91, 191, 232, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr := Get(adress)
	usr.Markets["0"] = 1000
	sell := trade.Sell{
		Adress:  []byte("0"),
		Offer:   1000,
		Recieve: 1000,
	}
	err := usr.AttachSell(sell)
	if err != nil {
		t.Error("this test should not throw any errors cuz its fine with user")
	}
	data.TestRM(adress)
}

func TestAttachSellWithNonExistingMarket(t *testing.T) {
	var adress = []byte{11, 122, 3, 44, 5, 16, 7, 8, 9, 110, 11, 12, 13, 14, 15, 16, 19, 18, 19, 20, 21, 22, 23, 24, 25, 232, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 138, 187, 40, 41, 42, 143, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 121, 59, 22, 91, 191, 232, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr := Get(adress)
	sell := trade.Sell{
		Adress:  []byte("0"),
		Offer:   1000,
		Recieve: 1000,
	}
	err := usr.AttachSell(sell)
	if err == nil {
		t.Error("this test should not throw error, cuz market should not exist")
	}
	data.TestRM(adress)
}
