package main

import (
	"fmt"
	"reflect"
	"sync_tree/_data"
)

func main() {
	testPutKey := []byte{1, 2, 3}
	testPutValue := []byte{1, 2, 3, 4, 5}
	testChangeValue := []byte{1, 2, 3, 4, 5, 6}
	_data.Put(testPutKey, testPutValue)
	fmt.Println("\033[32m(DATA) {Put} - passed\033[0m")

	checked := _data.Check(testPutKey)
	if checked {
		fmt.Println("\033[32m(DATA) {Check} - passed\033[0m")
	} else {
		fmt.Println("\033[31m(DATA) {Check} - failed\033[0m")
	}

	testGetValue := _data.Get(testPutKey)
	if reflect.DeepEqual(testGetValue, testPutValue) {
		fmt.Println("\033[32m(DATA) {Get} - passed\033[0m")
	} else {
		fmt.Println("\033[31m(DATA) {Get} - failed\033[0m")
	}

	_data.Change(testPutKey, testChangeValue)
	changed := _data.Get(testPutKey)
	if reflect.DeepEqual(changed, testChangeValue) {
		fmt.Println("\033[32m(DATA) {Change} - passed\033[0m")
	} else {
		fmt.Println("\033[32m(DATA) {Change} - failed\033[0m")
	}
}
