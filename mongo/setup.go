package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var database *mongo.Database
var userCollection *mongo.Collection

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

func createCollection(db *mongo.Database, name string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	err := db.CreateCollection(ctx, name)
	if err != nil {
		return err
	}
	return nil
}

func createIndex(coll *mongo.Collection, key string, value string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()
	idx := mongo.IndexModel{
		Keys: bsonx.Doc{{
			Key:   key,
			Value: bsonx.String(value),
		}},
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
		log.Panic("Unable to open mongo")
	}
	userCollectionError := createCollection(database, "user")
	if userCollectionError != nil {
		log.Panic("Unable to create user collection")
	}
	userCollection = database.Collection("user")
	userIndexError := createIndex(userCollection, "PubKey", "hash")
	if userIndexError != nil {
		log.Panic("Unable to create user hash PubKey index")
	}
}
