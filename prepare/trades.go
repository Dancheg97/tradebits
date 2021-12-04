package prepare

import (
	"sync_tree/calc"
	"sync_tree/market"
	"sync_tree/trade"
	"sync_tree/user"
)

func FullFillWithTrades() {
	btcMarket := market.Get(dummyMarketAdress1)

	firstDummyBuy := trade.Buy{
		Offer:   43000000,
		Recieve: 1800000,
	}
	secondDummyBuy := trade.Buy{
		Offer:   41500000,
		Recieve: 1790000,
	}
	thirdDummyBuy := trade.Buy{
		Offer:   41800000,
		Recieve: 1795000,
	}
	fourthDummyBuy := trade.Buy{
		Offer:   42100000,
		Recieve: 1895000,
	}
	fifthDummyBuy := trade.Buy{
		Offer:   42050000,
		Recieve: 1895398,
	}
	allDummyBuys := []trade.Buy{
		firstDummyBuy,
		secondDummyBuy,
		thirdDummyBuy,
		fourthDummyBuy,
		fifthDummyBuy,
	}
	for _, buy := range allDummyBuys {
		adr := calc.Rand()
		user.Create(adr, dummyMesKey, "dummy")
		usr := user.Get(adr)
		usr.Balance = 44800000
		usr.AttachBuy(&buy)
		btcMarket.AttachBuy(&buy)
	}

	firstDummySell := trade.Sell{
		Offer:   1799999,
		Recieve: 43000000,
	}
	secondDummySell := trade.Sell{
		Offer:   1799999,
		Recieve: 43000765,
	}
	thirdDummySell := trade.Sell{
		Offer:   1799999,
		Recieve: 43000853,
	}
	fourthDummySell := trade.Sell{
		Offer:   1799865,
		Recieve: 43000923,
	}
	fifthDummySell := trade.Sell{
		Offer:   1799212,
		Recieve: 43000999,
	}
	allDummySells := []trade.Sell{
		firstDummySell,
		secondDummySell,
		thirdDummySell,
		fourthDummySell,
		fifthDummySell,
	}
	for _, sell := range allDummySells {
		adr := calc.Rand()
		user.Create(adr, dummyMesKey, "dummy")
		usr := user.Get(adr)
		usr.Balances[string(dummyMarketAdress1)] = 44800000
		usr.AttachSell(&sell, dummyMarketAdress1)
		btcMarket.AttachSell(&sell)
	}
	btcMarket.Save()
}
