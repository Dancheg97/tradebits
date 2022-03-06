package mongo

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func openMongoFromEnv() error {
	godotenv.Load("../.env")
	mongo_host, _ := os.LookupEnv("MONGO_HOST")
	mongo_name, _ := os.LookupEnv("MONGO_NAME")
	mongo_password, _ := os.LookupEnv("MONGO_PASSWORD")
	mongo_db, _ := os.LookupEnv("MONGO_DB")
	return OpenMongo(
		mongo_host,
		mongo_name,
		mongo_password,
		mongo_db,
	)
}

func TestOpenMongo(t *testing.T) {
	err := openMongoFromEnv()
	if err != nil {
		t.Error("failed to open mongo")
	}
}

func TestCreateCollection(t *testing.T) {
	openMongoFromEnv()
	err := CreateCollection("testcol")
	if err != nil {
		t.Error("failed to create collection")
	}
	database.Collection("testcol").Drop(context.Background())
}

func TestCreateIndex(t *testing.T) {
	openMongoFromEnv()
	CreateCollection("testcol2")
	err := CreateIndex("testcol2", "Pubkey", "hashed")
	if err != nil {
		t.Error("failed to create collection")
	}
	database.Collection("testcol2").Drop(context.Background())
}
