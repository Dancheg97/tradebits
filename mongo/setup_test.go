package mongo

import (
	"context"
	"testing"
)

func TestOpenMongo(t *testing.T) {
	err := OpenMongo("mongodb://localhost:27017")
	if err != nil {
		t.Error("failed to open mongo")
	}
}

func TestCreateCollection(t *testing.T) {
	OpenMongo("mongodb://localhost:27017")
	err := CreateCollection("testcol")
	if err != nil {
		t.Error("failed to create collection")
	}
	database.Collection("testcol").Drop(context.Background())
}

func TestCreateIndex(t *testing.T) {
	OpenMongo("mongodb://localhost:27017")
	CreateCollection("testcol2")
	err := CreateIndex("testcol2", "Pubkey", "hashed")
	if err != nil {
		t.Error("failed to create collection")
	}
	database.Collection("testcol2").Drop(context.Background())
}
