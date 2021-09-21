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
		"There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc.",
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
		"There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc.",
		"https://image.flaticon.com/icons/png/512/1490/1490849.png",
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
}
