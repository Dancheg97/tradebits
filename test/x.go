package main

import (
	"path"
	"runtime"

	"github.com/syndtr/goleveldb/leveldb"
)

func openDB() *leveldb.DB {
	_, filename, _, _ := runtime.Caller(0)
	dbPath := path.Dir(filename) + "/leveldb"
	db, openErr := leveldb.OpenFile(dbPath, nil)
	if openErr != nil {
		return nil
	}
	return db
}

