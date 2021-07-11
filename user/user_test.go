package user

import (
	"sync_tree/data"
	"testing"
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
	var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 91, 91, 91, 91}
	var mesKey = []byte{1, 2, 3, 4, 5}
	var img = "user image link"
	Create()
	freeUser := Get(adress)
	defer freeUser.Save()
	if freeUser.PublicName == "user image link" {
		return
	}
	t.Error("get free user error")
}

// func TestGetBusyUser(t *testing.T) {
// 	freeUser := Get(adress)
// 	defer freeUser.Save()
// 	busyUser := Get(adress)
// 	if busyUser != nil {
// 		t.Error("attem to get free busy user succeded, test failed")
// 		return
// 	}
// }

// func TestChangeParameters(t *testing.T) {
// 	someUser := Get(adress)
// 	someUser.Balance = 1234
// 	someUser.Save()
// 	sameUser := Look(adress)
// 	if sameUser.Balance != 1234 {
// 		t.Error("same user balance should be 1000")
// 	}
// }
