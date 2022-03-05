package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"tradebits/mongo"
	"tradebits/redis"
	"tradebits/swagger"

	"github.com/joho/godotenv"
)

func readConfigField(field string) string {
	envar, found := os.LookupEnv(field)
	if !found {
		log.Fatal("problem loading .ENV field: ", envar)
	}
	return envar
}

func initMongo() {
	mongoErr := mongo.OpenMongo(
		readConfigField("MONGO_HOST"),
		readConfigField("MONGO_NAME"),
		readConfigField("MONGO_PASSWORD"),
		readConfigField("MONGO_DB"),
	)
	if mongoErr != nil {
		log.Fatal(mongoErr)
	}
	mongo.CreateCollection("user")
	mongo.CreateCollection("market")
	mongo.CreateCollection("trades")
	mongo.CreateCollection("network")
}

func initInfoResponse() {
	m := map[string]string{
		"name":      readConfigField("MARKET_NAME"),
		"mkey":      readConfigField("MARKET_PUBLICKEY"),
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
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	redis.Setup(readConfigField("REDIS_HOST"))
	initMongo()
	initInfoResponse()
}

func main() {
	log.Printf("Server starting...")
	router := swagger.NewRouter()
	log.Fatal(http.ListenAndServe(readConfigField("MARKET_PORT"), router))
}
