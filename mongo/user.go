package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	PubKey   string   `json:"PubKey" bson:"PubKey"`
	Balance  int      `json:"Balance" bson:"Balance"`
	Messages []string `json:"Messages" bson:"Messages"`
}

var userdb *mongo.Collection

func UserCheck(adress string) (bool, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	rez, err := userdb.Find(ctx, bson.M{
		"PubKey": adress,
	})
	if err != nil {
		return false, err
	}
	rezElems, err := rez.Current.Elements()
	if err != nil {
		return false, err
	}
	if len(rezElems) == 0 {
		return false, nil
	}
	return true, nil
}

func UserNew(u User) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	_, err := userdb.InsertOne(ctx, u)
	return err
}
