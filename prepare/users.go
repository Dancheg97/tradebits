package prepare

import (
	"encoding/pem"
	"io/ioutil"
	"strings"
	"sync_tree/calc"
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
