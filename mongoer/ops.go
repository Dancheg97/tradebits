package mongoer

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongoer) Check(key string, coll string) bool {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		45*time.Millisecond,
	)
	defer cancel()
	collection := m.database.Collection(coll)
	rez := collection.FindOne(ctx, bson.M{"key": key})
	return rez.Err() == nil
}

func (m *mongoer) Get(key string, coll string, i interface{}) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		45*time.Millisecond,
	)
	defer cancel()
	collection := m.database.Collection(coll)
	resp := collection.FindOne(ctx, bson.M{"key": key})
	if resp.Err() != nil {
		return resp.Err()
	}
	return resp.Decode(i)
}

func (m *mongoer) Put(coll string, i interface{}) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		45*time.Millisecond,
	)
	defer cancel()
	collection := m.database.Collection(coll)
	_, err := collection.InsertOne(ctx, i)
	return err
}

func (m *mongoer) Update(key string, coll string, i interface{}) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		45*time.Millisecond,
	)
	defer cancel()
	collection := m.database.Collection(coll)
	return collection.FindOneAndReplace(ctx, bson.M{"key": key}, i).Err()
}

func (m *mongoer) GetCollection(coll string) ([]map[string]string, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		45*time.Millisecond,
	)
	defer cancel()
	collection := m.database.Collection(coll)
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
