package mongoer

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongoer) Check(coll string, k string, v string) bool {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		45*time.Millisecond,
	)
	defer cancel()
	collection := m.database.Collection(coll)
	rez := collection.FindOne(ctx, bson.M{k: v})
	return rez.Err() == nil
}

func (m *mongoer) Get(coll string, k string, v string, i interface{}) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		45*time.Millisecond,
	)
	defer cancel()
	collection := m.database.Collection(coll)
	resp := collection.FindOne(ctx, bson.M{k: v})
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

func (m *mongoer) Update(coll string, k string, v string, i interface{}) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		45*time.Millisecond,
	)
	defer cancel()
	collection := m.database.Collection(coll)
	return collection.FindOneAndReplace(ctx, bson.M{k: v}, i).Err()
}

func (m *mongoer) GetCollection(coll string) ([]map[string]string, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		65*time.Millisecond,
	)
	defer cancel()
	collection := m.database.Collection(coll)
	results := []map[string]string{}
	cur, err := collection.Find(ctx, bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem map[string]string
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	return results, nil
}

func (m *mongoer) FindIdx(coll string, k string, v string) ([]string, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		65*time.Millisecond,
	)
	defer cancel()
	collection := m.database.Collection(coll)
	results := []string{}
	cur, err := collection.Find(ctx, bson.M{k: v}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem map[string]string
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem["_id"])
	}
	return results, nil
}

func (m *mongoer) GetIdx(coll string, id string, i interface{}) error {
	// TODO write when i wake up
	return nil
}
