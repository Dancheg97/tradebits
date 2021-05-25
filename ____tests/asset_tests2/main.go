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

func check_if_match(sell asset.Sell, buy asset.Buy) bool {
	return float64(sell.Offer/sell.Recieve) > float64(buy.Recieve/buy.Offer)
}

func check_full_close(sell asset.Sell, buy asset.Buy) bool {
	return buy.Offer >= sell.Recieve
}

// остаток sell Offer, новый buy
func full_close(sell asset.Sell, buy asset.Buy) (uint64, asset.Buy) {
	buy.Offer = buy.Offer - sell.Recieve
	buy.Recieve = sell.Recieve*
}

func main() {
	sell := asset.Sell{
		Offer:   900,
		Recieve: 100,
		Ratio:   900 / 100,
	}

	buy := asset.Buy{
		Offer:   300,
		Recieve: 2400,
		Ratio:   800 / 100,
	}

	fmt.Println(check_if_match(sell, buy))
	fmt.Println(check_full_close(sell, buy))
}
