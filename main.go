package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"tradebits/crypt"
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
	// TODO check case if mongo was already prepared
	// in case of container restart
	mongoErr := mongo.OpenMongo(
		readConfigField("MONGO_HOST"),
		readConfigField("MONGO_NAME"),
		readConfigField("MONGO_PASSWORD"),
		readConfigField("MONGO_DB"),
	)
	if mongoErr != nil {
		log.Fatal(mongoErr)
	}
	e1 := mongo.CreateCollection("net")
	e2 := mongo.CreateCollection("user")
	e3 := mongo.CreateCollection("market")
	e4 := mongo.CreateCollection("network")
	if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
		log.Fatal("Unable to set up Mongo database")
	}
	log.Println("mongo database set success")
}

func initInfoResponse() {
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

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	privKey := readConfigField("MARKET_PRIVATEKEY")
	keyErr := crypt.Setup(privKey)
	if keyErr != nil {
		log.Fatal(keyErr)
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
