package data

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var db, _ = leveldb.OpenFile("data/base", nil)

// get value by key from database
func Get(key []byte) []byte {
	output, getErr := db.Get(key, nil)
	if getErr != nil {
		return nil
	}
	return output
}

// put key by some value to database (if value exists use Change()
// func instead)
func Put(key []byte, value []byte) {
	valueExists, unexpected := db.Has(key, nil)
	if unexpected != nil {
		return
	}
	if valueExists {
		return
	}
	db.Put(key, value, nil)
}

// change existing value in database by key
func Change(key []byte, value []byte) {
	if Check(key) {
		db.Put(key, value, nil)
	}
}

// check if value exists in database
func Check(key []byte) bool {
	valueExists, unexpected := db.Has(key, nil)
	if unexpected != nil {
		return false
	}
	return valueExists
}
