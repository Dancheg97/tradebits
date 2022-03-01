package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// checks wether some value exists in database
func Check(key string, coll string) bool {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	collection := database.Collection(coll)
	rez := collection.FindOne(ctx, bson.M{"key": key})
	return rez.Err() == nil
}

// put some value from database to interface
func Get(key string, coll string, i interface{}) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	collection := database.Collection(coll)
	resp := collection.FindOne(ctx, bson.M{"key": key})
	if resp.Err() != nil {
		return resp.Err()
	}
	marshalErr := resp.Decode(i)
	if marshalErr != nil {
		return marshalErr
	}
	return nil
}

// puts some value to database
func Put(key string, coll string, i interface{}) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	collection := database.Collection(coll)
	_, err := collection.InsertOne(ctx, i)
	return err
}

// puts some value to database
func Update(key string, coll string, i interface{}) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	collection := database.Collection(coll)
	rez := collection.FindOneAndReplace(ctx, bson.M{"key": key}, i)
	return rez.Err()
}
