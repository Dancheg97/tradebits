package main

import (
	"fmt"
	"sync_tree/asset"
)

func checkMatch(sell asset.Sell, buy asset.Buy) bool {
	return float64(buy.Recieve/buy.Offer) < float64(sell.Offer/sell.Recieve)
}

func test_non_match() {
	sell := asset.Sell{
		Offer:   500,
		Recieve: 100,
	}
	buy := asset.Buy{
		Offer:   100,
		Recieve: 400,
	}
	if checkMatch(sell, buy) {
		fmt.Println("\033[32m[TEST] (ASSET_BUY) {MATCH} - passed\033[0m")
		return
	}
	fmt.Println("\033[31m[TEST] (ASSET_BUY) {MATCH} - failed\033[0m")
}

func main() {
	test_non_match()
}
