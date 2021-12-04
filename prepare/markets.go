package prepare

import (
	"sync_tree/calc"
	"sync_tree/market"
)

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
