package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Setup(adress string) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	credential := options.Credential{
		Username: "admin",
		Password: "admin",
	}
	opts := options.Client().ApplyURI(adress).SetAuth(credential)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Panic("Unable to connect to mongo")
	}
	db := client.Database("db")
	userdb = db.Collection("user")
}
