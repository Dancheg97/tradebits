package main

import (
	"fmt"
	"sync_tree/__logs"
	"sync_tree/user"
)

var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64}
var mesKey = []byte{1, 2, 3, 4, 5}
var img = "user image link"

func createNewUserTest() {
	creationErr := user.Create(adress, mesKey, img)
	if creationErr != nil {
		fmt.Println("\033[31m[TEST] (USER) {Create new} - failed\033[0m")
		return
	}
	fmt.Println("\033[32m[TEST] (USER) {Create new} - passed\033[0m")
}

func createExistingUserTest() {
	creationErr := user.Create(adress, mesKey, img)
	if creationErr != nil {
		fmt.Println("\033[32m[TEST] (USER) {Create existing} - passed\033[0m")
		return
	}
	fmt.Println("\033[31m[TEST] (USER) {Create existing} - failed\033[0m")
}

func getFreeUserTest() {
	freeUser := user.Get(adress)
	if freeUser.ImgLink == "user image link" {
		fmt.Println("\033[32m[TEST] (USER) {Get free} - passed\033[0m")
		return
	}
	fmt.Println("\033[31m[TEST] (USER) {Get free} - failed\033[0m")
}

func getBusyUserTest() {
	
}

func main() {
	__logs.Init()
	createNewUserTest()
	createExistingUserTest()
	getFreeUserTest()
}
