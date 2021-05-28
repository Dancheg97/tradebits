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
		__tests.Passed("asset", "CheckMatch", "checks if trade matches")
		return
	}
	__tests.Failed("asset", "CheckMatch", "checks if trade matches")
}

func main() {
	testRatioMatch()
}
