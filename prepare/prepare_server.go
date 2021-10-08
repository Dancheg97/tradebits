package prepare

import (
	"encoding/pem"
	"io/ioutil"
	"strings"
	"sync_tree/calc"
	"sync_tree/market"
	"sync_tree/trade"
	"sync_tree/user"
)

func stringToKeyBytes(key string) []byte {
	r := strings.NewReader(key)
	pemBytes, _ := ioutil.ReadAll(r)
	block, _ := pem.Decode(pemBytes)
	return block.Bytes
}

func CreateNewUsers() {
	alcoSplitted := strings.Split(alcoKeyString, "|")
	alcoAdress := calc.Hash(stringToKeyBytes(alcoSplitted[1]))
	alcoMesKey := stringToKeyBytes(alcoSplitted[3])
	alcErr := user.Create(alcoAdress, alcoMesKey, "Alcohol")
	if alcErr != nil {
		panic(alcErr)
	}
	alco := user.Get(alcoAdress)
	alco.Balance = 50000
	alco.Balances[string(dummyMarketAdress1)] = 10000
	alco.Save()

	nicoSplitted := strings.Split(nicoKeyString, "|")
	nicoAdress := calc.Hash(stringToKeyBytes(nicoSplitted[1]))
	nicoMesKey := stringToKeyBytes(nicoSplitted[3])
	nicErr := user.Create(nicoAdress, nicoMesKey, "Nicotin")
	if nicErr != nil {
		panic(nicErr)
	}
	nico := user.Get(nicoAdress)
	nico.Balance = 80000
	nico.Balances[string(dummyMarketAdress2)] = 4000
	nico.Save()
}

func CreateNewMarkets() {
	btcErr := market.Create(
		dummyMarketAdress1,
		"Bitcoin Ftem",
		dummyMesKey,
		dummyDescription,
		"https://image.flaticon.com/icons/png/512/1490/1490849.png",
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if btcErr != nil {
		panic(btcErr)
	}

	rubErr := market.Create(
		dummyMarketAdress2,
		"Sber ruble ftem",
		dummyMesKey,
		dummyDescription,
		"https://cdn-icons-png.flaticon.com/512/1548/1548946.png",
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if rubErr != nil {
		panic(rubErr)
	}

	cardErr := market.Create(
		calc.Rand(),
		"Cardano ftem",
		dummyMesKey,
		dummyDescription,
		"https://cdn-icons-png.flaticon.com/512/5245/5245441.png",
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if cardErr != nil {
		panic(cardErr)
	}

	zcashErr := market.Create(
		calc.Rand(),
		"Zcash ftem",
		dummyMesKey,
		dummyDescription,
		"https://cdn-icons-png.flaticon.com/512/1412/1412814.png",
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if zcashErr != nil {
		panic(zcashErr)
	}

	rippleErr := market.Create(
		calc.Rand(),
		"Ripple ftem",
		dummyMesKey,
		dummyDescription,
		"https://cdn-icons-png.flaticon.com/512/1181/1181387.png",
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if rippleErr != nil {
		panic(zcashErr)
	}

	dogeErr := market.Create(
		calc.Rand(),
		"Doge coin ftem",
		dummyMesKey,
		dummyDescription,
		"https://cdn-icons-png.flaticon.com/512/5004/5004807.png",
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if dogeErr != nil {
		panic(zcashErr)
	}

	ethErr := market.Create(
		calc.Rand(),
		"Ethereum ftem",
		dummyMesKey,
		dummyDescription,
		"https://cdn-icons-png.flaticon.com/512/1319/1319596.png",
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if ethErr != nil {
		panic(zcashErr)
	}
}

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
