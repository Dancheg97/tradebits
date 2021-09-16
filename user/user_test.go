package user

import (
	"reflect"
	"sync_tree/data"
	"sync_tree/trade"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 6, 5, 4, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 1, 3, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 2, 60, 61, 62, 91, 91}
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

func TestPutUserMessage(t *testing.T) {
	var adress = []byte{1, 22, 3, 44, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 1, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 32, 32, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 59, 56, 32, 121, 59, 22, 91, 191, 191, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr := Get(adress)
	usr.PutUserMessage([]byte{1, 2, 3}, "message")
	usr.Save()
	mes := Look(adress).GetMessages([]byte{1, 2, 3})[0]
	if mes != "umessage" {
		t.Error("the message should be 'message' - " + mes)
	}
	data.TestRM(adress)
}

func TestPutMarketMessage(t *testing.T) {
	var adress = []byte{1, 22, 3, 1, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 1, 23, 23, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 32, 32, 41, 42, 19, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 59, 56, 32, 121, 59, 22, 91, 191, 191, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr := Get(adress)
	usr.PutMarketMessage([]byte{1, 2, 3}, "message")
	usr.Save()
	mes := Look(adress).GetMessages([]byte{1, 2, 3})[0]
	if mes != "mmessage" {
		t.Error("the message should be 'message' - " + mes)
	}
	data.TestRM(adress)
}

<<<<<<< HEAD
=======
func TestNewUserNonNullableMessageMap(t *testing.T) {
	var adress = []byte{1, 22, 3, 1, 5, 6, 7, 8, 9, 123, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 1, 23, 23, 25, 26, 123, 123, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 32, 32, 41, 42, 19, 44, 45, 16, 47, 48, 49, 50, 51, 52, 53, 54, 59, 56, 32, 121, 59, 22, 91, 191, 191, 12}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr := Get(adress)
	if usr.Messages == nil {
		t.Error("user messages should never be null")
	}
	usr.Save()
	data.TestRM(adress)
}

>>>>>>> 4f25085bf170be3297aff057cd0611982d923edd
func TestAttachToLookedUser(t *testing.T) {
	var adress = []byte{1, 22, 3, 44, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 59, 56, 99, 121, 59, 22, 91, 191, 191, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
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
	data.TestRM(adress)
}

func TestAttachTradesWithZeroOffer(t *testing.T) {
	var adress = []byte{1, 22, 3, 44, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 19, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 121, 59, 22, 91, 91, 91, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
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
	data.TestRM(adress)
}

func TestAttachTradeWithBigBalance(t *testing.T) {
	var adress = []byte{1, 22, 3, 44, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 116, 19, 18, 19, 20, 21, 22, 23, 224, 25, 232, 27, 28, 29, 30, 31, 32, 33, 134, 35, 11, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 156, 57, 121, 59, 22, 91, 91, 232, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
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
	data.TestRM(adress)
}

func TestAttachNormalTrades(t *testing.T) {
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
	data.TestRM(adress)
}

func TestAttachSellNonExistingMarket(t *testing.T) {
	var adress = []byte{1, 22, 3, 44, 15, 6, 7, 8, 9, 110, 11, 12, 13, 14, 15, 16, 19, 18, 19, 121, 21, 22, 23, 24, 25, 232, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 138, 87, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 157, 121, 59, 22, 91, 91, 232, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
	usr := Get(adress)
	sell := trade.Sell{
		Offer:   1000,
		Recieve: 1000,
	}
	sellAttached := usr.AttachSell(&sell, []byte("x"))
	if sellAttached {
		t.Error("this sell should not be attached, cuz user dont have such market")
	}
	data.TestRM(adress)
}

func TestAttchBoundedTrades(t *testing.T) {
	var adress = []byte{1, 22, 3, 44, 22, 32, 7, 8, 9, 110, 11, 12, 13, 14, 15, 16, 19, 18, 19, 121, 21, 22, 23, 124, 25, 232, 27, 28, 29, 30, 31, 32, 33, 34, 35, 11, 37, 138, 87, 40, 41, 142, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 157, 121, 59, 122, 91, 91, 232, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create(adress, mesKey, img)
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
	data.TestRM(adress)
}
