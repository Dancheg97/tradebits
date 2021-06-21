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
	content, _ := ioutil.ReadFile("testKeys.pem")
	text := string(content)
	splitted := strings.Split(text, "|")
	key := splitted[1]
	r := strings.NewReader(key)
	pemBytes, _ := ioutil.ReadAll(r)
	block, _ := pem.Decode(pemBytes)
	adress := calc.Hash(block.Bytes)
	fmt.Println("putting to adress", adress)
	firstOne := user.Get(adress)
	if firstOne != nil {
		if firstOne.Balance == 0 {
			firstOne.Balance = 50000
		}
	}
	firstOne.Save()
}
