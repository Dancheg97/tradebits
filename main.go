package main

import (
	"log"
	"net/http"
	"os"

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

func init() {
	godotenv.Load()
	initCrypt()
	initResis()
	initMongo()
	initSwagger()
}

func main() {
	log.Printf("Server starting...")
	router := swagger.NewRouter()
	log.Fatal(http.ListenAndServe(readConfigField("MARKET_PORT"), router))
}
