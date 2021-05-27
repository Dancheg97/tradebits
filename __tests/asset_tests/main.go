package main

import (
	"reflect"
	"sync_tree/__tests"
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
		__tests.Failed("asset", "Create", "attempt to create new asset")
		return
	}
	__tests.Passed("asset", "Create", "attempt to create new asset")
}

func createExistingAssetTest() {
	err := asset.Create(adress, name, img, mesKey)
	if err != nil {
		__tests.Passed("asset", "Create", "attempt to create existing asset")
		return
	}
	__tests.Failed("asset", "Create", "attempt to create existing asset")
}

func getFreeAssetTest() {
	asset := asset.Get(adress)
	defer asset.Save()
	if reflect.DeepEqual(asset.MesKey, mesKey) {
		__tests.Passed("asset", "Get", "attempt to get free asset from db")
		return
	}
	__tests.Failed("asset", "Get", "attempt to get free asset from db")
}

func getBusyAssetTest() {
	freeAsset := asset.Get(adress)
	defer freeAsset.Save()
	busyAsset := asset.Get(adress)
	if busyAsset != nil {
		__tests.Failed("asset", "Get", "attempt to get busy asset from db")
		return
	}
	__tests.Passed("asset", "Get", "attempt to get busy asset from db")
}

func main() {
	__logs.Init()
	createNewTest()
	createExistingAssetTest()
	getFreeAssetTest()
	getBusyAssetTest()
}
