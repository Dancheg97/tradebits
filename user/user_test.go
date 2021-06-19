package user

import (
	"testing"
)

var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64}
var mesKey = []byte{1, 2, 3, 4, 5}
var img = "user image link"

func TestCreateUser(t *testing.T) {
	creationErr := Create(adress, mesKey, img)
	if creationErr != nil {
		t.Error("attemt to create new user failed")
		return
	}
}

func TestCreateExisting(t *testing.T) {
	creationErr := Create(adress, mesKey, img)
	if creationErr != nil {
		return
	}
	t.Error("attemt to create existing user succeded, error")
}

func TestGetFreeUser(t *testing.T) {
	freeUser := Get(adress)
	defer freeUser.Save()
	if freeUser.PublicName == "user image link" {
		return
	}
	t.Error("get free user error")
}

func TestGetBusyUser(t *testing.T) {
	freeUser := Get(adress)
	defer freeUser.Save()
	busyUser := Get(adress)
	if busyUser != nil {
		t.Error("attem to get free busy user succeded, test failed")
		return
	}
}
