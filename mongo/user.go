package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

type user struct {
	Pubkey   string   `bson:"Pubkey"`
	Balance  int      `bson:"Balance"`
	Messages []string `bson:"Messages"`
}

func User() *user {
	
}

// checks wether some value exists in database
func UserCheck(pubkey string) bool {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	rez := userCollection.FindOne(ctx, bson.M{"Pubkey": pubkey})
	if rez.Err() != nil {
		return false
	}
	return true
}

// create new user in db
func UserCreate(pubkey string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	_, err := userCollection.InsertOne(ctx, bson.M{
		"Pubkey":   pubkey,
		"Balance":  0,
		"Messages": []string{},
	})
	return err
}

// save some value to database
func UserGet(pubkey string) (*user, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	resp := userCollection.FindOne(ctx, bson.M{"Pubkey": pubkey})
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	u := user{}
	marshalErr := resp.Decode(&u)
	if marshalErr != nil {
		return nil, marshalErr
	}
	return &u, nil
}

// update user in database
func UserSave(u *user) 