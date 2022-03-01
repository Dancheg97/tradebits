package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database

func openMongo(adress string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	credential := options.Credential{
		Username: "admin",
		Password: "admin",
	}
	opts := options.Client().ApplyURI(adress).SetAuth(credential)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}
	database = client.Database("main")
	return nil
}

func createCollection(name string) error {
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

func createIndex(col string, key string, value string) error {
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

func Setup(adress string) {
	dbOpenError := openMongo(adress)
	if dbOpenError != nil {
		log.Panic("Unable to open mongo", dbOpenError)
	}
	userCollectionError := createCollection("user")
	if userCollectionError != nil {
		log.Panic("Unable to create user collection", userCollectionError)
	}
	userCollection = database.Collection("user")
	userIndexError := createIndex("user", "PubKey", "hashed")
	if userIndexError != nil {
		log.Panic("Unable to create user hash PubKey index", userIndexError)
	}
}
