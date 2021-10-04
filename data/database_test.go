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

func TestChangeValue(t *testing.T) {
	key := []byte{1, 1, 8, 3}
	firstVal := []byte{1, 1, 8, 8}
	secondVal := []byte{1, 1, 8, 9}
	Put(key, firstVal)
	Change(key, secondVal)
	getVal := Get(key)
	if reflect.DeepEqual(getVal, firstVal) {
		t.Error("the get value should be equal to second")
	}
	base.Delete(key, nil)
}

func TestCheckExisting(t *testing.T) {
	key := []byte{1, 1, 8, 12}
	val := []byte{1, 1, 8, 12}
	Put(key, val)
	exists := Check(key)
	base.Delete(key, nil)
	if !exists {
		t.Error("value should exist")
	}
	base.Delete(key, nil)
}

func TestCheckNotExisting(t *testing.T) {
	key := []byte{1, 1, 8, 16}
	exists := Check(key)
	if exists {
		t.Error("this value should not exist")
	}
}

func TestRemoveValue(t *testing.T) {
	val := []byte{1, 1, 8, 16}
	Put(val, val)
	TestRM(val)
	exists := Check(val)
	if exists {
		t.Error("this value should not exists, cuz it have been removed")
	}
}
