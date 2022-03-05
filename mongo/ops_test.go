package mongo

import (
	"context"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTestEnv(collectioname string) {
	openMongoFromEnv()
	CreateCollection(collectioname)
	CreateIndex(collectioname, "key", "hashed")
	go func() {
		time.Sleep(time.Millisecond * 100)
		database.Collection(collectioname).Drop(context.Background())
	}()
}

func TestPut(t *testing.T) {
	collectionname := "testputcol"
	setupTestEnv(collectionname)
	err := Put(collectionname, &map[string]string{
		"key": "testputkey",
	})
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Millisecond * 200)
}

func TestCheck(t *testing.T) {
	collectionname := "testcheckcol"
	setupTestEnv(collectionname)
	Put(collectionname, &map[string]string{
		"key": "testcheckkey",
	})
	found := Check("testcheckkey", collectionname)
	if !found {
		t.Error("not found")
	}
	time.Sleep(time.Millisecond * 200)
}

func TestGet(t *testing.T) {
	collectionname := "testgetcol"
	setupTestEnv(collectionname)
	Put(collectionname, &map[string]string{
		"key":   "testgetkey",
		"vaval": "tester",
	})
	mp := map[string]string{}
	Get("testgetkey", collectionname, &mp)
	if mp["vaval"] != "tester" {
		t.Error("mongo returned bad value")
	}
	time.Sleep(time.Millisecond * 200)
}

func TestUpdate(t *testing.T) {
	collectionname := "testupdatecol"
	setupTestEnv(collectionname)
	Put(collectionname, &map[string]string{
		"key":   "testupdatekey",
		"vaval": "tester",
	})
	rez := Update("testupdatekey", collectionname, &map[string]string{
		"key":   "testupdatekey",
		"vaval": "tester2",
	})
	if rez != nil {
		t.Error("Error with updating")
	}
	mp := map[string]string{}
	Get("testupdatekey", collectionname, &mp)
	if mp["vaval"] != "tester2" {
		t.Error("mongo did not update value in database")
	}
}

func TestGetCollection(t *testing.T) {
	collectionname := "testgetwholecollection"
	setupTestEnv(collectionname)
	Put(collectionname, &map[string]string{
		"vaval": "tester1",
	})
	Put(collectionname, &map[string]string{
		"vaval": "tester2",
	})
	results := []map[string]string{}
	cur, err := database.Collection(collectionname).Find(
		context.TODO(), bson.D{{}}, options.Find(),
	)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem map[string]string
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	t.Error(results)
}
