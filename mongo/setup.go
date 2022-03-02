package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database

func OpenMongo(adress string, user string, pass string, db string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	credential := options.Credential{
		Username: user,
		Password: pass,
	}
	opts := options.Client().ApplyURI(adress).SetAuth(credential)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}
	database = client.Database(db)
	return nil
}

func CreateCollection(name string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	err := database.CreateCollection(ctx, name)
	if err != nil {
		return err
	}
	return nil
}

func CreateIndex(col string, key string, value string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	coll := database.Collection(col)
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
