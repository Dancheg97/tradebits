package main

import (
	"log"
	"os"
)

var file, _ = os.OpenFile("logs/info.log", os.O_CREATE|os.O_APPEND, 0644)

func main() {
	log.Print("error")
	

}