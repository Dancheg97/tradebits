package _data

import (
	"sync_tree/__logs"

	"github.com/syndtr/goleveldb/leveldb"
)

var db, _ = leveldb.OpenFile("_data/base", nil)

// get value by key from database
func Get(key []byte) []byte {
	output, getErr := db.Get(key, nil)
	if getErr != nil {
		__logs.Critical("unexpected error Get non existant value")
		return nil
	}
	return output
}

// put key by some value to database (if value exists use Change()
// func instead)
func Put(key []byte, value []byte) {
	valueExists, unexpected := db.Has(key, nil)
	if unexpected != nil {
		__logs.Critical("unexpected error in db on Put func")
		return
	}
	if valueExists {
		__logs.Critical("value exists and shouldn't be changed")
		return
	}
	db.Put(key, value, nil)
}

// change existing value in database by key
func Change(key []byte, value []byte) {
	if Check(key) {

	}
	dbErr := db.Put(key, value, nil)
	if dbErr != nil {
		__logs.Critical("unexpected error in db on Change func")
		return
	}
}

// check if value exists in database
func Check(key []byte) bool {
	valueExists, unexpected := db.Has(key, nil)
	if unexpected != nil {
		__logs.Critical("unexpected error in db on Check func")
		return false
	}
	return valueExists
}
