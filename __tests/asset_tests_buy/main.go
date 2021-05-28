package main

import (
	"sync_tree/__tests"
	"sync_tree/asset"
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
	if asset.B2SCheckMatch(sell, buy) {
		__tests.Passed("asset", "Buy/Match", "checks if trade matches")
		return
	}
	__tests.Failed("asset", "Buy/Match", "checks if trade matches")
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
	if asset.B2SCheckMatch(sell, buy) {
		__tests.Failed("asset", "Buy/-Match", "checks if trade not matches")
		return
	}
	__tests.Passed("asset", "Buy/-Match", "checks if trade not matches")
}

func testIfSellIsClosing() {
	sell := asset.Sell{
		OfferAsset:  500,
		RecieveMain: 100,
	}
	buy := asset.Buy{
		OfferMain:    80,
		RecieveAsset: 200,
	}
	if asset.B2SCheckMatch(sell, buy) {
		if asset.B2SIfCloseSeller(sell, buy) {
			__tests.Failed("asset", "Buy/Sell-", "checks if trade closes sell")
			return
		}
	}
	__tests.Passed("asset", "Buy/Sell-", "checks if trade closes sell")
}

func testIfBuyIsClosing() {
	sell := asset.Sell{
		OfferAsset:  500,
		RecieveMain: 100,
	}
	buy := asset.Buy{
		OfferMain:    800,
		RecieveAsset: 2000,
	}
	if asset.B2SCheckMatch(sell, buy) {
		if asset.B2SIfCloseSeller(sell, buy) {
			__tests.Passed("asset", "Buy/Buy-", "checks if trade closes buy")
			return
		}
	}
	__tests.Failed("asset", "Buy/Buy-", "checks if trade closes buy")
}

func testCloseSell() {
	sell := asset.Sell{
		OfferAsset:  500,
		RecieveMain: 100,
	}
	buy := asset.Buy{
		OfferMain:    120,
		RecieveAsset: 200,
	}
	if asset.B2SCheckMatch(sell, buy) {
		if asset.B2SIfCloseSeller(sell, buy) {
			newBuy, sellOut, buyOut := asset.B2SCloseSeller(sell, buy)
			ch1 := sellOut.MainOut == 100
			ch2 := sellOut.AssetOut == 333
			ch3 := buyOut.AssetOut == 167
			ch4 := newBuy.OfferMain == 20
			ch5 := newBuy.RecieveAsset == 33
			if ch1 && ch2 && ch3 && ch4 && ch5 {
				__tests.Passed("asset", "CloseSeller", "pretty good numbers")
				return
			}
			__tests.Failed("asset", "CloseSeller", "bad numbers")
			return
		} else {
			__tests.Failed("asset", "CloseSeller", "not closing")
			return
		}
	}
	__tests.Failed("asset", "CloseSeller", "not even matching")
}

func testCloseBuyer() {
	sell := asset.Sell{
		OfferAsset:  500,
		RecieveMain: 100,
	}
	buy := asset.Buy{
		OfferMain:    80,
		RecieveAsset: 200,
	}
	if asset.B2SCheckMatch(sell, buy) {
		if asset.B2SIfCloseSeller(sell, buy) {
			__tests.Failed("asset", "CloseBuyer", "not closing")
			return
		} else {
			newSell, sellOut, buyOut := asset.B2SCloseBuyer(sell, buy)
			ch1 := sellOut.MainOut == 80
			ch2 := buyOut.AssetOut == 200
			ch3 := newSell.OfferAsset == 300
			ch4 := newSell.RecieveMain == 20
			if ch1 && ch2 && ch3 && ch4 {
				__tests.Passed("asset", "CloseBuyer", "pretty good numbers")
				return
			}
			__tests.Failed("asset", "CloseBuyer", "bad numbers")
			return
		}
	}
	__tests.Failed("asset", "CloseBuyer", "not even matching")
}

func main() {
	testMatch()
	testNonMatch()
	testIfSellIsClosing()
	testIfBuyIsClosing()
	testCloseSell()
	testCloseBuyer()
}
