package database

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var DB, _ = leveldb.OpenFile("database/data", nil)
