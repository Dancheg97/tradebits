package data

import (
	"reflect"
	"testing"
)

func TestPutValue(t *testing.T) {
	key := []byte{1, 2, 3}
	val := []byte{1, 2, 3, 4, 5}
	Put(key, val)
	val2 := Get(key)
	if !reflect.DeepEqual(val, val2) {
		t.Error("values are not equal")
	}
	TestRM(key)
}

func TestPutExisting(t *testing.T) {
	key := []byte{1, 2, 2}
	val := []byte{1, 2, 2, 4, 5}
	Put(key, val)
	val2 := []byte{1, 2, 2, 4, 5, 5}
	Put(key, val2)
	val3 := Get(key)
	if reflect.DeepEqual(val3, val2) {
		t.Error("value in db should not change")
	}
	base.Delete(key, nil)
}
