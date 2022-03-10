package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Server starting...")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(readConfigField("MARKET_PORT"), router))
}
