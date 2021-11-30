package data

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

type dataObject struct {
	key   []byte
	value []byte
}

var reserveNodeConnected = false
var DataQue = []dataObject{}

func GetIterator() (iterator.Iterator, *leveldb.DB) {
	reserveNodeConnected = true
	base.Close()
	db := openDB()
	iterator := db.NewIterator(nil, nil)
	return iterator, db
}

func ReOpenBase(db *leveldb.DB) []dataObject {
	db.Close()
	base = openDB()
	return DataQue
}

func CloseConnection() {
	DataQue = []dataObject{}
	reserveNodeConnected = false
}
