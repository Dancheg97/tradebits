package main

import (
	"fmt"
	"sync_tree/asset"
)

func checkMatch(sell asset.Sell, buy asset.Buy) bool {
	return float64(buy.Recieve/buy.Offer) < float64(sell.Offer/sell.Recieve)
}

func ifCloseBuyer(sell asset.Sell, buy asset.Buy) bool {
	return sell.Recieve > buy.Offer
}

func closeBuyer(sell asset.Sell, buy asset.Buy) (asset.Sell )  {

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

func testIfBuyerIsClosing() {
	sell := asset.Sell{
		Offer:   500,
		Recieve: 100,
	}
	buy := asset.Buy{
		Offer:   80,
		Recieve: 200,
	}
	if checkMatch(sell, buy) {
		if ifCloseBuyer(sell, buy) {
			fmt.Println("\033[32m(ASSET_BUY) {REST_SELLER} - passed\033[0m")
			return
		}
	}
	fmt.Println("\033[31m(ASSET_BUY) {REST_SELLER} - failed\033[0m")
}

func testIfSellerIsClosing() {
	sell := asset.Sell{
		Offer:   500,
		Recieve: 100,
	}
	buy := asset.Buy{
		Offer:   800,
		Recieve: 2000,
	}
	if checkMatch(sell, buy) {
		if ifCloseBuyer(sell, buy) {
			fmt.Println("\033[31m(ASSET_BUY) {REST_BUYER} - failed\033[0m")
			return
		}
	}
	fmt.Println("\033[32m(ASSET_BUY) {REST_BUYER} - passed\033[0m")
}

func testCloseSell() {
	sell := asset.Sell{
		Offer:   500, // B
		Recieve: 100, // A
	}
	buy := asset.Buy{
		Offer:   120, // A
		Recieve: 200, // B 
	}
	if checkMatch(sell, buy) {
		if ifCloseBuyer(sell, buy) {

		}
	}
}

func closeSeller(sell asset.Sell, buy asset.Buy) (asset.Sell, asset.Buy) {
	goesToSeller = uint64(buy.Recieve/buy.Offer*sell.Recieve)
	leavesToBuyer = buy.Recieve - goesToSeller

}

func main() {
	testMatch()
	testNonMatch()
	testRestSeller()
	testRestBuyer()
}
