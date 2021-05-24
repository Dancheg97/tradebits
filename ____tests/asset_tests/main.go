package main

import (
	"fmt"
	"reflect"
	"sync_tree/__logs"
	"sync_tree/asset"
)

var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 63}
var mesKey = []byte{1, 2, 3, 4, 5}
var img = "asset image link . example"
var name = "newAsset"

func createNewTest() {
	err := asset.Create(adress, name, img, mesKey)
	if err != nil {
		fmt.Println("\033[31m[TEST] (ASSET) {Create new} - failed\033[0m")
		return
	}
	fmt.Println("\033[32m[TEST] (ASSET) {Create new} - passed\033[0m")
}

func createExistingAssetTest() {
	creationErr := asset.Create(adress, name, img, mesKey)
	if creationErr != nil {
		fmt.Println("\033[32m[TEST] (ASSET) {Create existing} - passed\033[0m")
		return
	}
	fmt.Println("\033[31m[TEST] (ASSET) {Create existing} - failed\033[0m")
}

func getFreeAssetTest() {
	asset := asset.Get(adress)
	defer asset.Save()
	if reflect.DeepEqual(asset.MesKey, mesKey) {
		fmt.Println("\033[32m[TEST] (ASSET) {Get} - passed\033[0m")
		return
	}
	fmt.Println("\033[31m[TEST] (ASSET) {Get} - failed\033[0m")
}

func getBusyAssetTest() {
	freeAsset := asset.Get(adress)
	defer freeAsset.Save()
	busyAsset := asset.Get(adress)
	if busyAsset != nil {
		fmt.Println("\033[31m[TEST] (ASSET) {Get busy} - failed\033[0m")
		return
	}
	fmt.Println("\033[32m[TEST] (ASSET) {Get busy} - passed\033[0m")
}

func main() {
	__logs.Init()
	createNewTest()
	createExistingAssetTest()
	getFreeAssetTest()
	getBusyAssetTest()
}
