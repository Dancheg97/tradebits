package database

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var db, _ = leveldb.OpenFile("./data", nil)
