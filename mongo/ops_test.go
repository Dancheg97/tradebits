package mongo

import (
	"context"
	"testing"
	"time"
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
		t.Error("retrieved bad value")
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
}
