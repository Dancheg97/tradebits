package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"tradebits/api"
	"tradebits/crypter"
	"tradebits/mongoer"
	"tradebits/rediser"

	"github.com/joho/godotenv"
)

func readConfigField(field string) string {
	envar, found := os.LookupEnv(field)
	if !found {
		log.Fatal("problem loading .ENV field: ", envar)
	}
	return envar
}

func init() {
	godotenv.Load()
	readConfigField("REDIS_HOST")
	crypt, err1 := crypter.Get(readConfigField("MARKET_PRIVATEKEY"))
	redis, err2 := rediser.Get(readConfigField("REDIS_HOST"))
	mongo, err3 := mongoer.Get(
		readConfigField("MONGO_HOST"),
		readConfigField("MONGO_NAME"),
		readConfigField("MONGO_PASSWORD"),
		readConfigField("MONGO_DB"),
	)
	m := map[string]string{
		"name":      readConfigField("MARKET_NAME"),
		"mkey":      crypt.Pub(),
		"descr":     readConfigField("MARKET_DESCR"),
		"img":       readConfigField("MARKET_IMG"),
		"worktime":  readConfigField("MARKET_WORKTIME"),
		"fee":       readConfigField("MARKET_FEE"),
		"delimiter": readConfigField("MARKET_DELIMITER"),
	}
	respbytes, err4 := json.Marshal(m)
	if !(err1 != nil && err2 != nil && err3 != nil && err4 != nil) {
		log.Fatal("Setup api error: ", err1, err2, err3, err4)
	}
	api.Setup(respbytes, mongo, crypt, redis)
}

func main() {
	log.Printf("Server starting...")
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(readConfigField("MARKET_PORT"), router))
}
