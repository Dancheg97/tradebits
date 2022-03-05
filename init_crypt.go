package main

import (
	"log"
	"tradebits/crypt"
)

func initCrypt() {
	privKey := readConfigField("MARKET_PRIVATEKEY")
	err := crypt.Setup(privKey)
	if err != nil {
		log.Fatal("Setup crypt error: ", err)
	}
	log.Println("Setup crypt success")
}
