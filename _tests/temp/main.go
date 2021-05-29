package main

import (
	"fmt"
	"reflect"
	"sync_tree/data"
)

func main() {
	testPutKey := []byte{1, 2, 3}
	testPutValue := []byte{1, 2, 3, 4, 5}
	testChangeValue := []byte{1, 2, 3, 4, 5, 6}
	data.Put(testPutKey, testPutValue)

	checked := data.Check(testPutKey)
	if checked {
	} else {
		fmt.Println("failed to check item existens in database")
	}

	testGetValue := data.Get(testPutKey)
	if reflect.DeepEqual(testGetValue, testPutValue) {
	} else {
		fmt.Println("failed to get item from database")
	}

	data.Change(testPutKey, testChangeValue)
	changed := data.Get(testPutKey)
	if reflect.DeepEqual(changed, testChangeValue) {
	} else {
		fmt.Println("failed to change item contents in database")
	}
}
