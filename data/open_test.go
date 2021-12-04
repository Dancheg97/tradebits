package data

import (
	"testing"
)

func TestRemoveValue(t *testing.T) {
	val := []byte{1, 1, 8, 16}
	Put(val, val)
	TestRM(val)
	exists := Check(val)
	if exists {
		t.Error("this value should not exists, cuz it have been removed")
	}
}
