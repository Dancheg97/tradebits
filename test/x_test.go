package main

import (
	"testing"
)

func TestBaseIterator(t *testing.T) {
	base := openDB()
	base.Put([]byte{0}, []byte{0}, nil)
	base.Put([]byte{1}, []byte{1}, nil)
	base.Put([]byte{2}, []byte{2}, nil)
	itr := base.NewIterator(nil, nil)
	base.Put()
	itr.Next()
	t.Error(itr.Key(), itr.Value())
	itr.Next()
	t.Error(itr.Key(), itr.Value())
	itr.Next()
	t.Error(itr.Key(), itr.Value())
	t.Error(itr.Next())
}
