package main

import (
	"log"
	"tradebits/mongo"
)

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
