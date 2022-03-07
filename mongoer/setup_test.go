package mongoer

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func getTestMongoer() (*mongoer, error) {
	godotenv.Load("../.env")
	mongo_host, _ := os.LookupEnv("MONGO_HOST")
	return Get(mongo_host)
}

func TestOpenMongo(t *testing.T) {
	_, err := getTestMongoer()
	if err != nil {
		t.Error("failed to open mongo")
	}
}

func TestCreateCollection(t *testing.T) {
	m, _ := getTestMongoer()
	err := m.CreateCollection("testcol")
	if err != nil {
		t.Error("failed to create collection")
	}
	m.database.Collection("testcol").Drop(context.Background())
}

func TestCreateIndex(t *testing.T) {
	m, _ := getTestMongoer()
	m.CreateCollection("testcol2")
	err := m.CreateIndex("testcol2", "Pubkey", "hashed")
	if err != nil {
		t.Error("failed to create collection")
	}
	m.database.Collection("testcol2").Drop(context.Background())
}
