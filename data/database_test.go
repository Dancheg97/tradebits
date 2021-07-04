package data

import (
	"reflect"
	"testing"
)

func TestDataBase(t *testing.T) {
	testPutKey := []byte{1, 2, 3}
	testPutValue := []byte{1, 2, 3, 4, 5}
	testChangeValue := []byte{1, 2, 3, 4, 5, 6}
	Put(testPutKey, testPutValue)

	checked := Check(testPutKey)
	if checked {
	} else {
		t.Error("failed to check item existens in database")
	}

	testGetValue := Get(testPutKey)
	if reflect.DeepEqual(testGetValue, testPutValue) {
	} else {
		t.Error("failed to get item from database")
	}

	Change(testPutKey, testChangeValue)
	changed := Get(testPutKey)
	if reflect.DeepEqual(changed, testChangeValue) {
	} else {
		t.Error("failed to change item contents in database")
	}
}
