package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	PubKey   string   `json:"PubKey" bson:"PubKey"`
	Balance  int      `json:"Balance" bson:"Balance"`
	Messages []string `json:"Messages" bson:"Messages"`
}

var userdb *mongo.Collection

func NewUser(u User) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	_, err := userdb.InsertOne(ctx, u)
	return err
}
