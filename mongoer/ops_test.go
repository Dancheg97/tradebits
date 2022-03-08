package mongoer

import (
	"context"
	"testing"
	"time"
)

func setupTestEnv(collectioname string) *mongoer {
	m, _ := getTestMongoer()
	m.CreateCollection(collectioname)
	m.CreateIndex(collectioname, "key", "hashed")
	go func() {
		time.Sleep(time.Millisecond * 100)
		m.database.Collection(collectioname).Drop(context.Background())
	}()
	return m
}

func TestPut(t *testing.T) {
	collectionname := "testputcol"
	m := setupTestEnv(collectionname)
	err := m.Put(collectionname, &map[string]string{
		"key": "testputkey",
	})
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Millisecond * 200)
}

func TestCheck(t *testing.T) {
	collectionname := "testcheckcol"
	m := setupTestEnv(collectionname)
	m.Put(collectionname, &map[string]string{
		"key": "testcheckkey",
	})
	found := m.Check("testcheckkey", collectionname)
	if !found {
		t.Error("not found")
	}
	time.Sleep(time.Millisecond * 200)
}

func TestGet(t *testing.T) {
	collectionname := "testgetcol"
	m := setupTestEnv(collectionname)
	m.Put(collectionname, &map[string]string{
		"key":   "testgetkey",
		"vaval": "tester",
	})
	mp := map[string]string{}
	m.Get("testgetkey", collectionname, &mp)
	if mp["vaval"] != "tester" {
		t.Error("mongo returned bad value")
	}
	time.Sleep(time.Millisecond * 200)
}

func TestUpdate(t *testing.T) {
	collectionname := "testupdatecol"
	m := setupTestEnv(collectionname)
	m.Put(collectionname, &map[string]string{
		"key":   "testupdatekey",
		"vaval": "tester",
	})
	rez := m.Update("testupdatekey", collectionname, &map[string]string{
		"key":   "testupdatekey",
		"vaval": "tester2",
	})
	if rez != nil {
		t.Error("Error with updating")
	}
	mp := map[string]string{}
	m.Get("testupdatekey", collectionname, &mp)
	if mp["vaval"] != "tester2" {
		t.Error("mongo did not update value in database")
	}
}

func TestGetCollection(t *testing.T) {
	collectionname := "testgetwholecollection"
	m := setupTestEnv(collectionname)
	m.Put(collectionname, &map[string]string{
		"vaval": "tester1",
	})
	m.Put(collectionname, &map[string]string{
		"vaval": "tester2",
	})
	vals, err := m.GetCollection(collectionname)
	if err != nil {
		t.Error(err)
	}
	if len(vals) < 1 {
		t.Error("bad length of output collection elements")
	}
	time.Sleep(time.Millisecond * 200)
}

func TestFindIndexes(t *testing.T) {
	collectionname := "testgetindexes"
	m := setupTestEnv(collectionname)
	m.CreateIndex(collectionname, "vaval", "hashed")
	m.Put(collectionname, &map[string]string{
		"vaval": "tester",
	})
	m.Put(collectionname, &map[string]string{
		"vaval": "tester",
	})
	vals, err := m.FindIndexes(collectionname, "vaval", "tester")
	if err != nil {
		t.Error(err)
	}
	if len(vals) < 1 {
		t.Error("bad length of output collection elements")
	}
	time.Sleep(time.Millisecond * 200)
}
