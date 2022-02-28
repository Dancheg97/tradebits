package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"Name" bson:"Name"`
	Balance  int                `json:"Balance" bson:"Balance"`
	Messages []string           `json:"Messages" bson:"Messages"`
}

var client *mongo.Client

func Setup(adress string) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	mongoclient, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(adress),
	)
	client = mongoclient
	if err != nil {
		log.Panic("Unable to connect to mongo")
	}
}
