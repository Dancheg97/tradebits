package main

import (
	"fmt"
	"sync_tree/asset"
)

func checkMatch(sell asset.Sell, buy asset.Buy) bool {
	return float64(buy.Recieve/buy.Offer) < float64(sell.Offer/sell.Recieve)
}

func restSeller(sell asset.Sell, buy asset.Buy) bool {
	return sell.Recieve > buy.Offer
}

func testMatch() {
	sell := asset.Sell{
		Offer:   500,
		Recieve: 100,
	}
	buy := asset.Buy{
		Offer:   100,
		Recieve: 400,
	}
	if checkMatch(sell, buy) {
		fmt.Println("\033[32m(ASSET_BUY) {MATCH} - passed\033[0m")
		return
	}
	fmt.Println("\033[31m(ASSET_BUY) {MATCH} - failed\033[0m")
}

func testNonMatch() {
	sell := asset.Sell{
		Offer:   500,
		Recieve: 100,
	}
	buy := asset.Buy{
		Offer:   100,
		Recieve: 600,
	}
	if checkMatch(sell, buy) {
		fmt.Println("\033[31m(ASSET_BUY) {NON_MATCH} - failed\033[0m")
		return
	}
	fmt.Println("\033[32m(ASSET_BUY) {NON_MATCH} - passed\033[0m")
}

func testCloseSeller() {
	sell := asset.Sell{
		Offer:   500,
		Recieve: 100,
	}
	buy := asset.Buy{
		Offer:   80,
		Recieve: 200,
	}
	if checkMatch(sell, buy) {
		if restSeller(sell, buy) {
			fmt.Println("\033[32m(ASSET_BUY) {REST_SELLER} - passed\033[0m")
			return
		}
	}
	fmt.Println("\033[31m(ASSET_BUY) {REST_SELLER} - failed\033[0m")
}

func testCloseBuyer() {
	sell := asset.Sell{
		Offer:   500,
		Recieve: 100,
	}
	buy := asset.Buy{
		Offer:   800,
		Recieve: 2000,
	}
	if checkMatch(sell, buy) {
		if restSeller(sell, buy) {
			fmt.Println("\033[31m(ASSET_BUY) {REST_BUYER} - failed\033[0m")
			return
		}
	}
	fmt.Println("\033[32m(ASSET_BUY) {REST_BUYER} - passed\033[0m")
}

func main() {
	testMatch()
	testNonMatch()
	testCloseSeller()
	testCloseBuyer()
}
