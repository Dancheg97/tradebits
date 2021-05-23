package _data

import (
	"sync_tree/__logs"

	"github.com/syndtr/goleveldb/leveldb"
)

var db, _ = leveldb.OpenFile("_data/base", nil)

func Get(key []byte) []byte {
	output, getErr := db.Get(key, nil)
	if getErr != nil {
		__logs.Critical("unexpected error Get non existant value")
		return nil
	}
	return output
}

func Put(key []byte, value []byte) {
	valueExists, unexpected := db.Has(key, nil)
	if unexpected != nil {
		__logs.Critical("unexpected error in db on Put func")
		return
	}
	if valueExists {
		__logs.Critical("unexpected error in db on Put func")
		return
	}
	db.Put(key, value, nil)
}

func Change(key []byte, value []byte) {
	dbErr := db.Put(key, value, nil)
	if dbErr != nil {
		__logs.Critical("unexpected error in db on Change func")
		return
	}
}

func Check(key []byte) bool {
	valueExists, unexpected := db.Has(key, nil)
	if unexpected != nil {
		__logs.Critical("unexpected error in db on Check func")
		return false
	}
	return valueExists
}
