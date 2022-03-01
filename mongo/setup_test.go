package mongo

import (
	"context"
	"testing"
)

func TestOpenMongo(t *testing.T) {
	err := openMongo("mongodb://localhost:27017")
	if err != nil {
		t.Error("failed to open mongo")
	}
}

func TestCreateCollection(t *testing.T) {
	openMongo("mongodb://localhost:27017")
	err := createCollection("testcol")
	if err != nil {
		t.Error("failed to create collection")
	}
	database.Collection("testcol").Drop(context.Background())
}

func TestCreateIndex(t *testing.T) {
	openMongo("mongodb://localhost:27017")
	createCollection("testcol2")
	err := createIndex("testcol2", "Pubkey", "hashed")
	if err != nil {
		t.Error("failed to create collection")
	}
	database.Collection("testcol2").Drop(context.Background())
}

func TestSetup(t *testing.T) {
	Setup("mongodb://localhost:27017")
	database.Collection("user").Drop(context.Background())
}
