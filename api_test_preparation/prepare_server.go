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
	user.Create(alcoAdress, alcoMesKey, "Alcohol")
	alco := user.Get(alcoAdress)
	alco.Balance = 50000
	alco.Balances[string(alcoAdress)] = 10000
	alco.Save()

	nicoSplitted := strings.Split(nicoKeyString, "|")
	nicoAdress := calc.Hash(stringToKeyBytes(nicoSplitted[1]))
	nicoMesKey := stringToKeyBytes(nicoSplitted[3])
	user.Create(nicoAdress, nicoMesKey, "Nicotin")
	nico := user.Get(nicoAdress)
	nico.Balance = 80000
	nico.Balances[string(nicoAdress)] = 4000
	nico.Save()
}

func CreateNewMarkets() {
	alcoSplitted := strings.Split(alcoKeyString, "|")
	alcoAdress := calc.Hash(stringToKeyBytes(alcoSplitted[1]))
	market.Create(
		alcoAdress,
		"Bitcoin Ftem",
		dummyMessageKey,
		dummyDescription,
		"https://image.flaticon.com/icons/png/512/1490/1490849.png",
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)

	nicoSplitted := strings.Split(nicoKeyString, "|")
	nicoAdress := calc.Hash(stringToKeyBytes(nicoSplitted[1]))
	market.Create(
		nicoAdress,
		"Sber ruble ftem",
		dummyMessageKey,
		dummyDescription,
		"https://cdn-icons-png.flaticon.com/512/1548/1548946.png",
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)

	market.Create(
		calc.Rand(),
		"Cardano",
		dummyMessageKey,
		dummyDescription,
		"https://cdn-icons-png.flaticon.com/512/5245/5245441.png",
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
}
