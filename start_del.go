package main

import (
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync_tree/calc"
	"sync_tree/market"
	"sync_tree/user"

	"github.com/blevesearch/bleve/v2"
)

func stringToKeyBytes(key string) []byte {
	r := strings.NewReader(key)
	pemBytes, _ := ioutil.ReadAll(r)
	block, _ := pem.Decode(pemBytes)
	return block.Bytes
}

func createNewUsers() {
	alcoContent, _ := ioutil.ReadFile("Alcohol.pem")
	alcoText := string(alcoContent)
	alcoSplitted := strings.Split(alcoText, "|")
	alcoAdress := calc.Hash(stringToKeyBytes(alcoSplitted[1]))
	alcoMesKey := stringToKeyBytes(alcoSplitted[3])
	user.Create(alcoAdress, alcoMesKey, "Alcohol")
	alco := user.Get(alcoAdress)
	defer alco.Save()
	if alco.Balance == 0 {
		alco.Balance = 50000
	}
	fmt.Println("alco wallet created with", alco.Balance, "balance")
	nicoContent, _ := ioutil.ReadFile("Nicotin.pem")
	nicoText := string(nicoContent)
	nicoSplitted := strings.Split(nicoText, "|")
	nicoAdress := calc.Hash(stringToKeyBytes(nicoSplitted[1]))
	nicoMesKey := stringToKeyBytes(nicoSplitted[3])
	user.Create(nicoAdress, nicoMesKey, "Nicotin")
	nico := user.Get(nicoAdress)
	defer nico.Save()
	if nico.Balance == 0 {
		nico.Balance = 50000
	}
	fmt.Println("nico wallet created with", nico.Balance, "balance")
}

func createStartMarket() {
	if _, err := os.Stat("data/search"); os.IsNotExist(err) {
		mapping := bleve.NewIndexMapping()
		bleve.New("data/search", mapping)
	}
	market.Create(
		[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 63},
		"bitcoin first market",
		[]byte{0, 1, 2, 3, 4, 2, 8},
		"First market to be created on a platform, just for testing",
		"https://image.flaticon.com/icons/png/512/1490/1490849.png",
	)
	market.Create(
		[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64},
		"ruble second market",
		[]byte{0, 1, 2, 3, 4, 2, 8},
		"Second market to be created on a platform also for testing",
		"https://image.flaticon.com/icons/png/512/1490/1490839.png",
	)
}
