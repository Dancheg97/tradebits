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
	respbytes, infoEjectErr := json.Marshal(m)
	if infoEjectErr != nil {
		log.Panic("failed to marshall info to bytes")
	}
	swagger.MarketInfoResponse = respbytes
	log.Println("info response set success")
}
