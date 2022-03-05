package main

import (
	"log"
	"tradebits/crypt"
)

func initCrypt() {
	privKey := readConfigField("MARKET_PRIVATEKEY")
	keyErr := crypt.Setup(privKey)
	if keyErr != nil {
		log.Fatal(keyErr)
	}
}
