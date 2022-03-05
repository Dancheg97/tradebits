package main

import (
	"encoding/json"
	"log"
	"tradebits/crypt"
	"tradebits/swagger"
)

func initSwagger() {
	m := map[string]string{
		"name":      readConfigField("MARKET_NAME"),
		"mkey":      crypt.Pub,
		"descr":     readConfigField("MARKET_DESCR"),
		"img":       readConfigField("MARKET_IMG"),
		"worktime":  readConfigField("MARKET_WORKTIME"),
		"fee":       readConfigField("MARKET_FEE"),
		"delimiter": readConfigField("MARKET_DELIMITER"),
	}
	respbytes, err := json.Marshal(m)
	if err != nil {
		log.Fatal("Setup swagger error: ", err)
	}
	swagger.MarketInfoResponse = respbytes
	log.Println("Setup swagger success")
}