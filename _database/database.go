package database

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var db, _ = leveldb.OpenFile("database/userData", nil)

type Data interface {
	Put(key []byte, value []byte)
	Get(key []byte) []byte
}

