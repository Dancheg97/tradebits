package mongo

import (
	"context"
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

func Setup() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	client, _ := mongo.Connect(ctx, &options.ClientOptions{})
}
