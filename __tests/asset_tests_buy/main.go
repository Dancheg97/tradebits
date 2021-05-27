package main

import (
	"fmt"
	"sync_tree/asset"
	"sync_tree/____tests"
)

func testMatch() {
	sell := asset.Sell{
		OfferAsset:  500,
		RecieveMain: 100,
	}
	buy := asset.Buy{
		OfferMain:    100,
		RecieveAsset: 400,
	}
	if asset.CheckMatch(sell, buy) {
		____tests.Passed("asset", "CheckMatch", "checks if asset matches")
		return
	}
	fmt.Println("\033[31m(ASSET_BUY) {MATCH} - failed\033[0m")
}

func testNonMatch() {
	sell := asset.Sell{
		OfferAsset:  500,
		RecieveMain: 100,
	}
	buy := asset.Buy{
		OfferMain:    100,
		RecieveAsset: 600,
	}
	if asset.CheckMatch(sell, buy) {
		fmt.Println("\033[31m(ASSET_BUY) {NON_MATCH} - failed\033[0m")
		return
	}
	fmt.Println("\033[32m(ASSET_BUY) {NON_MATCH} - passed\033[0m")
}

func testIfBuyerIsClosing() {
	sell := asset.Sell{
		OfferAsset:  500,
		RecieveMain: 100,
	}
	buy := asset.Buy{
		OfferMain:    80,
		RecieveAsset: 200,
	}
	if asset.CheckMatch(sell, buy) {
		if asset.IfCloseBuyer(sell, buy) {
			fmt.Println("\033[32m(ASSET_BUY) {REST_SELLER} - passed\033[0m")
			return
		}
	}
	fmt.Println("\033[31m(ASSET_BUY) {REST_SELLER} - failed\033[0m")
}

func testIfSellerIsClosing() {
	sell := asset.Sell{
		OfferAsset:  500,
		RecieveMain: 100,
	}
	buy := asset.Buy{
		OfferMain:    800,
		RecieveAsset: 2000,
	}
	if asset.CheckMatch(sell, buy) {
		if asset.IfCloseBuyer(sell, buy) {
			fmt.Println("\033[31m(ASSET_BUY) {REST_BUYER} - failed\033[0m")
			return
		}
	}
	fmt.Println("\033[32m(ASSET_BUY) {REST_BUYER} - passed\033[0m")
}

func testCloseSell() {
	sell := asset.Sell{
		OfferAsset:  500,
		RecieveMain: 100,
	}
	buy := asset.Buy{
		OfferMain:    80,
		RecieveAsset: 200,
	}
	if asset.CheckMatch(sell, buy) {
		if asset.IfCloseBuyer(sell, buy) {
			fmt.Println("\033[31m(ASSET_BUY) {REST_BUYER} - failed\033[0m")
			return
		} else {
			newBuy, sellOut, buyOut := asset.CloseSeller(sell, buy)
			fmt.Println("new buy request", newBuy)
			fmt.Println("output for seller", sellOut)
			fmt.Println("output for buyer", buyOut)

		}
	}
}

func main() {
	testMatch()
	testNonMatch()
	testIfBuyerIsClosing()
	testIfSellerIsClosing()
	testCloseSell()
}
