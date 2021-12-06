package main

import (
	"testing"
)

func TestBaseIterator(t *testing.T) {
	base := openDB()
	go base.Put([]byte{0}, []byte{0}, nil)
	go base.Put([]byte{1}, []byte{1}, nil)
	go base.Put([]byte{2}, []byte{2}, nil)
	itr := base.NewIterator(nil, nil)
	go base.Put([]byte{0}, []byte{4}, nil)
	go base.Put([]byte{1}, []byte{4}, nil)
	go base.Put([]byte{2}, []byte{4}, nil)
	go base.Put([]byte{3}, []byte{4}, nil)
	itr.Next()
	t.Error(itr.Key(), itr.Value())
	itr.Next()
	t.Error(itr.Key(), itr.Value())
	itr.Next()
	t.Error(itr.Key(), itr.Value())
	t.Error(itr.Next())
}
