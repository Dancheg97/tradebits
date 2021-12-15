package data

import (
	"reflect"
	"testing"
)

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
