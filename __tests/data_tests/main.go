package main

import (
	"reflect"
	"sync_tree/_data"
	"sync_tree/__tests"
)

func main() {
	testPutKey := []byte{1, 2, 3}
	testPutValue := []byte{1, 2, 3, 4, 5}
	testChangeValue := []byte{1, 2, 3, 4, 5, 6}
	_data.Put(testPutKey, testPutValue)
	__tests.Passed("data", "Put", "put item to database")

	checked := _data.Check(testPutKey)
	if checked {
		__tests.Passed("data", "Check", "put item to database")
	} else {
		__tests.Failed("data", "Check", "put item to database")
	}

	testGetValue := _data.Get(testPutKey)
	if reflect.DeepEqual(testGetValue, testPutValue) {
		__tests.Passed("data", "Get", "get item from database")
	} else {
		__tests.Failed("data", "Get", "get item from database")
	}

	_data.Change(testPutKey, testChangeValue)
	changed := _data.Get(testPutKey)
	if reflect.DeepEqual(changed, testChangeValue) {
		__tests.Passed("data", "Change", "change item value in database")
	} else {
		__tests.Failed("data", "Change", "change item value in database")
	}
}
