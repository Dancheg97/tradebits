package data

import (
	"path"
	"runtime"

	"github.com/syndtr/goleveldb/leveldb"
)

var base = openDB()

func openDB() *leveldb.DB {
	_, filename, _, _ := runtime.Caller(0)
	dbPath := path.Dir(filename) + "/leveldb"
	db, openErr := leveldb.OpenFile(dbPath, nil)
	if openErr != nil {
		return nil
	}
	return db
}

// get value by key from database
func Get(key []byte) []byte {
	output, getErr := base.Get(key, nil)
	if getErr != nil {
		return nil
	}
	return output
}

// put key by some value to database (if value exists use Change()
// func instead)
func Put(key []byte, value []byte) {
	valueExists, _ := base.Has(key, nil)
	if valueExists {
		return
	}
	base.Put(key, value, nil)
}

// change existing value in database by key
func Change(key []byte, value []byte) {
	if Check(key) {
		base.Put(key, value, nil)
	}
}

// check if value exists in database
func Check(key []byte) bool {
	valueExists, _ := base.Has(key, nil)
	return valueExists
}

// function is made only to remove values after testing, dont call it in
// any other case
func TestRM(key []byte) {
	base.Delete(key, nil)
}
