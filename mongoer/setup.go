package mongoer

import (
	"context"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongoer interface {
	CreateCollection(name string) error
	CreateIndex(col string, key string, value string) error
	Check(key string, coll string) bool
	Get(key string, coll string, i interface{}) error
	Put(coll string, i interface{}) error
	Update(key string, coll string, i interface{}) error
	GetCollection(coll string) ([]map[string]string, error)
}

type mongoer struct {
	database *mongo.Database
}

func Get(adr string) (*mongoer, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		45*time.Millisecond,
	)
	defer cancel()
	opts := options.Client().ApplyURI(adr)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	m := mongoer{
		database: client.Database("main"),
	}
	return &m, nil
}

func (m *mongoer) CreateCollection(name string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		45*time.Millisecond,
	)
	defer cancel()
	err := m.database.CreateCollection(ctx, name)
	if err != nil {
		if strings.Contains(err.Error(), "Collection already exists") {
			return nil
		}
		return err
	}
	return nil
}

func (m *mongoer) CreateIndex(col string, key string, value string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		45*time.Millisecond,
	)
	defer cancel()
	coll := m.database.Collection(col)
	idx := mongo.IndexModel{
		Keys:    bson.M{key: value},
		Options: nil,
	}
	_, err := coll.Indexes().CreateOne(ctx, idx)
	if err != nil {
		return err
	}
	return nil
}
