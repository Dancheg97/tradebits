package data

import "testing"

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
