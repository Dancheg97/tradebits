package api_test_preparation

import (
	"encoding/pem"
	"io/ioutil"
	"strings"
	"sync_tree/calc"
	"sync_tree/market"
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
}
