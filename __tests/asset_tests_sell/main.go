package main

import (
	"sync_tree/__tests"
	"sync_tree/asset"
)

func testRatioMatch() {
	buy := asset.Buy{
		OfferMain:    500,
		RecieveAsset: 200,
	}
	sell := asset.Sell{
		OfferAsset:  500,
		RecieveMain: 100,
	}
	if asset.B2SCheckMatch(buy, sell) {
		__tests.Passed("asset", "CheckMatch", "checks if trade matches")
	}
}

func main() {
	testRatioMatch()
}
