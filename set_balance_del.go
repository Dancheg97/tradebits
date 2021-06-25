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
		firstOne.Balance = 50000
	}
	firstOne.Save()
}

func lookAtBalance() {
	content, _ := ioutil.ReadFile("Alcohol.pem")
	text := string(content)
	splitted := strings.Split(text, "|")
	key := splitted[1]
	r := strings.NewReader(key)
	pemBytes, _ := ioutil.ReadAll(r)
	block, _ := pem.Decode(pemBytes)
	adress := calc.Hash(block.Bytes)
	firstOne := user.Look(adress)
	fmt.Println(firstOne)
}
