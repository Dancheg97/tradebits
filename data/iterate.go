package data

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

func GetIterator() (iterator.Iterator, *leveldb.DB) {
	base.Close()
	db := openDB()
	iterator := db.NewIterator(nil, nil)
	return iterator, db
}

func ReturnIterator(itr iterator.Iterator, db *leveldb.DB) {
	db.Close()
	base = openDB()
}
