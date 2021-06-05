package data

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var leafs, _ = leveldb.OpenFile("data/leafs", nil)
var branches, _ = leveldb.OpenFile("data/branches", nil)

// get value by key from database
func Get(key []byte) []byte {
	output, getErr := leafs.Get(key, nil)
	if getErr != nil {
		return nil
	}
	return output
}

// put key by some value to database (if value exists use Change()
// func instead)
func Put(key []byte, value []byte) {
	valueExists, unexpected := leafs.Has(key, nil)
	if unexpected != nil {
		return
	}
	if valueExists {
		return
	}
	leafs.Put(key, value, nil)
}

// change existing value in database by key
func Change(key []byte, value []byte) {
	if Check(key) {
		leafs.Put(key, value, nil)
	}
}

// check if value exists in database
func Check(key []byte) bool {
	valueExists, unexpected := leafs.Has(key, nil)
	if unexpected != nil {
		return false
	}
	return valueExists
}

// write transaction to database
func WriteTransaction(hash []byte, content []byte) {
	branches.Put(hash, content, nil)
}
