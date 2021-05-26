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

func check_if_sell_match(sell asset.Sell, buy asset.Buy) bool {
	return float64(sell.Offer/sell.Recieve) > float64(buy.Recieve/buy.Offer)
}

func check_full_close(sell asset.Sell, buy asset.Buy) bool {
	return buy.Offer >= sell.Recieve
}

// остаток sell Offer, новый buy
func full_close(sell asset.Sell, buy asset.Buy) (uint64, asset.Buy) {
	buyRatio := float64(buy.Recieve / buy.Offer)
	buy.Offer = buy.Offer - sell.Recieve
	newRecieve := buy.Recieve - uint64(float64(sell.Recieve)*buyRatio)
	rest := sell.Offer - uint64(float64(sell.Recieve)*buyRatio)
	buy.Recieve = newRecieve
	return rest, buy
}

func match(sell asset.Sell, buy asset.Buy) (asset.Sell, asset.Buy) {
	buyRatio := 
}

func main() {
	/* 
	всего есть 3 кейса:
	1 - сделка не состоится
	2 - сделка состоится частично для покупателя
	3 - сделка состоится частично для продавца
	*/
	sell := asset.Sell{
		Offer:   900,
		Recieve: 100,
	}

	buy := asset.Buy{
		Offer:   200,
		Recieve: 856,
	}
	if check_if_sell_match(sell, buy) {
		fmt.Println(check_if_sell_match(sell, buy))
		if check_full_close(sell, buy) {
			fmt.Println(check_full_close(sell, buy))
			fmt.Println(full_close(sell, buy))
		}
	}

}
