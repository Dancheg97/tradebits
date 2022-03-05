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
	return resp.Decode(i)
}

// puts some value to database
func Put(coll string, i interface{}) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	collection := database.Collection(coll)
	_, err := collection.InsertOne(ctx, i)
	return err
}

// changes some value in database
func Update(key string, coll string, i interface{}) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	collection := database.Collection(coll)
	return collection.FindOneAndReplace(ctx, bson.M{"key": key}, i).Err()
}

func GetCollection(coll string, i interface{}) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	collection := database.Collection(coll)
	idx, err := collection.Indexes().List(ctx)
	if err != nil {
		return err
	}

	for {
		id := idx.ID()
		resp := collection.FindOne(ctx, bson.M{"_id": id})
		if resp.Err() != nil {
			return resp.Err()
		}
		decodeErr := resp.Decode(i)
		if decodeErr != nil {
			return decodeErr
		}
	}

}
