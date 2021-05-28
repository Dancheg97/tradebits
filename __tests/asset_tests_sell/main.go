package main

import (
	"sync_tree/__tests"
	"sync_tree/asset"
)

func testRatioMatch() {
	buy := asset.Buy{
		OfferMain:    400,
		RecieveAsset: 200,
	}
	sell := asset.Sell{
		OfferAsset:  400,
		RecieveMain: 200,
	}
	if asset.B2SCheckMatch(buy, sell) {
		__tests.Passed("asset", "Sell/Match", "checks if trade matches")
		return
	}
	__tests.Failed("asset", "Sell/Match", "checks if trade matches")
}

func testRatioNonMatch() {
	buy := asset.Buy{
		OfferMain:    100,
		RecieveAsset: 200,
	}
	sell := asset.Sell{
		OfferAsset:  100,
		RecieveMain: 200,
	}
	if asset.B2SCheckMatch(buy, sell) {
		__tests.Failed("asset", "Sell/Match", "check if trade not matches")
	}
	__tests.Passed("asset", "Sell/Match", "check if trade not matches")
}

func testCloseSeller() {
	
}

func main() {
	testRatioMatch()
	testRatioNonMatch()
}
