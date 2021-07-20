package main

import (
	"encoding/pem"
	"fmt"
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

func main() {
	marketAdress := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 63}
	mkt := market.Get(marketAdress)
	fmt.Println("mkt")
	fmt.Println(mkt)
	alcoContent, _ := ioutil.ReadFile("Alcohol.pem")
	alcoText := string(alcoContent)
	alcoSplitted := strings.Split(alcoText, "|")
	alcoAdress := calc.Hash(stringToKeyBytes(alcoSplitted[1]))
	alco := user.Get(alcoAdress)
	fmt.Println("alco")
	fmt.Println(alco)
	nicoContent, _ := ioutil.ReadFile("Nicotin.pem")
	nicoText := string(nicoContent)
	nicoSplitted := strings.Split(nicoText, "|")
	nicoAdress := calc.Hash(stringToKeyBytes(nicoSplitted[1]))
	nico := user.Get(nicoAdress)
	fmt.Println("nico")
	fmt.Println(nico)
}
