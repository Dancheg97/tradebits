package main

import (
	"fmt"
	"sync_tree/asset"
)

func operateSellToBuy(sell asset.Sell, buy asset.Buy) (asset.Sell, asset.Buy) {
	if float64(sell.Offer/sell.Recieve) < float64(buy.Recieve/buy.Offer) {
		return sell, buy
	}
	if buy.Offer < sell.Recieve {
		sell.Offer = sell.Offer - buy.Recieve
		sell.Recieve = sell.Recieve - buy.Offer
		buy.Offer = 0
		buy.Recieve = 0
		return sell, buy
	}
	newSellerOffer := uint64(buy.Recieve / buy.Offer *
		(buy.Offer - sell.Recieve))
	newBuyOffer := buy.Offer - sell.Recieve
	newBuyRecieve := sell.Offer + buy.Recieve
	sell.Offer = newSellerOffer
	sell.Recieve = 0
	buy.Offer = newBuyOffer
	buy.Recieve = newBuyRecieve
	return sell, buy
}

func main() {
	sell := asset.Sell{
		Offer:   500,
		Recieve: 100,
	}

	buy := asset.Buy{
		Offer:   50,
		Recieve: 5000,
	}

	fmt.Println(operateSellToBuy(sell, buy))
}
