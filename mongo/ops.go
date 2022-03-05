package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetCollection(coll string) ([]map[string]string, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	collection := database.Collection(coll)
	results := []map[string]string{}
	cur, err := collection.Find(ctx, bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var elem map[string]string
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	return results, nil
}
