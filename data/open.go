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

// function is made only to remove values after generating some
// test data, dont call it in any other case
func TestRM(key []byte) {
	base.Delete(key, nil)
}
