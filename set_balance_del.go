package main

import (
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"strings"
	"sync_tree/calc"
	"sync_tree/user"
)

func setStartBalance() {
	// dont forget that this section works only in case user with requered
	// adress is already created
	content, _ := ioutil.ReadFile("Alcohol.pem")
	text := string(content)
	splitted := strings.Split(text, "|")
	key := splitted[1]
	r := strings.NewReader(key)
	pemBytes, _ := ioutil.ReadAll(r)
	block, _ := pem.Decode(pemBytes)
	adress := calc.Hash(block.Bytes)
	firstOne := user.Get(adress)
	if firstOne.Balance == 0 {
		fmt.Println("yes, first user balance was 0, changing that to 100.000")
		firstOne.Balance = 100000
	}
	firstOne.Save()
}
