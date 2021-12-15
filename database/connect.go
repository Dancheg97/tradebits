package database

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

type databaseObject struct {
	key   []byte
	value []byte
}

var reserveNodeConnected = false
var databaseQue = []databaseObject{}

func GetIterator() (iterator.Iterator, *leveldb.DB) {
	reserveNodeConnected = true
	base.Close()
	db := openDB()
	iterator := db.NewIterator(nil, nil)
	return iterator, db
}

func ReOpenBase(db *leveldb.DB) []databaseObject {
	db.Close()
	base = openDB()
	return databaseQue
}

func CloseConnection() {
	databaseQue = []databaseObject{}
	reserveNodeConnected = false
}
