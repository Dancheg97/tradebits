package main

import (
	"sync_tree/__tests"
	"sync_tree/__logs"
	"sync_tree/user"
)

var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64}
var mesKey = []byte{1, 2, 3, 4, 5}
var img = "user image link"

func createNewUserTest() {
	creationErr := user.Create(adress, mesKey, img)
	if creationErr != nil {
		__tests.Failed("user", "Create", "attempt to create new user")
		return
	}
	__tests.Passed("user", "Create", "attempt to create new user")
}

func createExistingUserTest() {
	creationErr := user.Create(adress, mesKey, img)
	if creationErr != nil {
		__tests.Passed("user", "Create", "attempt to create existing user")
		return
	}
	__tests.Failed("user", "Create", "attempt to create existing user")
}

func getFreeUserTest() {
	freeUser := user.Get(adress)
	defer freeUser.Save()
	if freeUser.ImgLink == "user image link" {
		__tests.Passed("user", "Get", "attempt get free user")
		return
	}
	__tests.Failed("user", "Get", "attempt get free user")
}

func getBusyUserTest() {
	freeUser := user.Get(adress)
	defer freeUser.Save()
	busyUser := user.Get(adress)
	if busyUser != nil {
		__tests.Failed("user", "Create", "attempt get free user")
		return
	}
	__tests.Passed("user", "Create", "attempt get free user")
}

func main() {
	__logs.Init()
	createNewUserTest()
	createExistingUserTest()
	getFreeUserTest()
	getBusyUserTest()
}
