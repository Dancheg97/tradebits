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
	err := m.Put(collectionname, &map[string]string{"k": "v"})
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Millisecond * 200)
}

func TestCheck(t *testing.T) {
	collectionname := "testcheckcol"
	m := setupTestEnv(collectionname)
	m.Put(collectionname, &map[string]string{"k": "v"})
	found := m.Check(collectionname, "k", "v")
	if !found {
		t.Error("not found")
	}
	time.Sleep(time.Millisecond * 200)
}

func TestGet(t *testing.T) {
	collectionname := "testgetcol"
	m := setupTestEnv(collectionname)
	m.Put(collectionname, &map[string]string{"k": "v", "k2": "v2"})
	mp := map[string]string{}
	m.Get(collectionname, "k", "v", &mp)
	if mp["k2"] != "v2" {
		t.Error("mongo returned incorrect value")
	}
	time.Sleep(time.Millisecond * 200)
}

func TestUpdate(t *testing.T) {
	collectionname := "testupdatecol"
	m := setupTestEnv(collectionname)
	m.Put(collectionname, &map[string]string{"k": "v", "k2": "v2"})
	rez := m.Update(
		collectionname,
		"k", "v",
		&map[string]string{"k": "v", "k2": "v3"},
	)
	if rez != nil {
		t.Error("Error with updating")
	}
	mp := map[string]string{}
	m.Get(collectionname, "k", "v", &mp)
	if mp["k2"] != "v3" {
		t.Error("mongo did not update value in database")
	}
}

func TestGetCollection(t *testing.T) {
	collectionname := "testgetwholecollection"
	m := setupTestEnv(collectionname)
	m.Put(collectionname, &map[string]string{"k": "v"})
	m.Put(collectionname, &map[string]string{"k": "v"})
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
	m.CreateIndex(collectionname, "k", "hashed")
	m.Put(collectionname, &map[string]string{"k": "v"})
	m.Put(collectionname, &map[string]string{"k": "v"})
	vals, err := m.FindIdx(collectionname, "k", "v")
	if err != nil {
		t.Error(err)
	}
	if len(vals) < 1 {
		t.Error("bad length of output collection elements")
	}
	time.Sleep(time.Millisecond * 200)
}

func TestGetIdx(t *testing.T) {
	collectionname := "testgetindexes"
	m := setupTestEnv(collectionname)
	m.CreateIndex(collectionname, "k", "hashed")
	m.Put(collectionname, &map[string]string{"k": "v"})
	vals, _ := m.FindIdx(collectionname, "k", "v")
	mp := map[string]string{}
	err := m.GetIdx(collectionname, vals[0], &mp)
	if err != nil {
		t.Error(err)
	}
	if mp["k"] != "v" {
		t.Error("value not matching")
	}
	time.Sleep(time.Millisecond * 200)
}

func TestGet2kv(t *testing.T) {
	collectionname := "testgetindexes"
	m := setupTestEnv(collectionname)
	m.CreateIndex(collectionname, "k", "hashed")
	m.Put(collectionname, &map[string]string{"k1": "v1", "k2": "v2"})
	m.Put(collectionname, &map[string]string{"k1": "v2", "k2": "v1"})
	mp := map[string]string{}
	err := m.Get2kv(collectionname, "k1", "v1", "k2", "v2", &mp)
	if err != nil {
		t.Error(err)
	}
	if mp["k1"] != "v1" || mp["k2"] != "v2" {
		t.Error("value not matching")
	}
	time.Sleep(time.Millisecond * 200)
}
