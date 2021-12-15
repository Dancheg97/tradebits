package database

import (
	"reflect"
	"testing"
)

func TestGetValue(t *testing.T) {
	key := []byte{1, 1, 2}
	val := []byte{1, 1, 1}
	Put(key, val)
	val2 := Get(key)
	if !reflect.DeepEqual(val, val2) {
		t.Error("values should be the same")
	}
	base.Delete(key, nil)
}

func TestGetNotExisting(t *testing.T) {
	key := []byte{1, 1, 8}
	val := Get(key)
	if val != nil {
		t.Error("this key does not exist")
	}
}
