package main

import (
	"log"
	"tradebits/redis"
)

func initResis() {
	err := redis.Setup(readConfigField("REDIS_HOST"))
	if err != nil {
		log.Fatal("Setup redis error: ", err)
	}
	log.Println("Setup redis success")
}
