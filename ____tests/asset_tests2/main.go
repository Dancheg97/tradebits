package main

import (
	"fmt"
	"sync_tree/asset"
)

/*
make trade
make existed
make with close
make with open
*/

func sellMatches(sell asset.Sell, buy asset.Buy) bool {
	return float64(sell.Offer/sell.Recieve) > float64(buy.Recieve/buy.Offer)
}

func fullForSeller(sell asset.Sell, buy asset.Buy) bool {
	return buy.Offer >= sell.Recieve
}

func closeSeller(sell asset.Sell, buy asset.Buy) (asset.Sell, asset.Buy) {
	newSellerOffer := uint64(buy.Recieve / buy.Offer *
		(buy.Offer - sell.Recieve))
	newBuyOffer := buy.Offer - sell.Recieve
	newBuyRecieve := sell.Offer - newSellerOffer
	sell.Offer = newSellerOffer
	sell.Recieve = 0
	buy.Offer = newBuyOffer
	buy.Recieve = newBuyRecieve
	return sell, buy
}

func closeBuyer(sell asset.Sell, buy asset.Buy) (asset.Sell, asset.Buy) {
	sell.Offer = sell.Offer - buy.Recieve
	sell.Recieve = sell.Recieve - buy.Offer
	buy.Offer = 0
	buy.Recieve = 0
	return sell, buy
}

func main() {
	sell := asset.Sell{
		Offer:   900,
		Recieve: 100,
	}

	buy := asset.Buy{
		Offer:   50,
		Recieve: 400,
	}

	if sellMatches(sell, buy) {
		if fullForSeller(sell, buy) {
			fmt.Println(closeSeller(sell, buy))
		} else {
			fmt.Println(closeBuyer(sell, buy))
		}
	}

}
