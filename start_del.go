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
		[]byte{0, 1, 2, 3, 4, 4, 6},
		"bitcoin and other shit twice",
		[]byte{0, 1, 2, 3, 4, 2, 8},
		"First market to be created on a platform",
		"https://image.flaticon.com/icons/png/512/1490/1490849.png",
	)
}
