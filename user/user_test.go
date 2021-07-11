package user

import (
	"reflect"
	"sync_tree/data"
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

}
