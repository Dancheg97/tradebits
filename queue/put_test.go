package queue

import (
	"reflect"
	"testing"
)

func TestPutSomeMessages(t *testing.T) {
	newque := Create()
	firstVal := KvPair{
		Key:   []byte{0, 1, 2},
		Value: []byte{1, 2, 3},
	}
	secondVal := KvPair{
		Key:   []byte{0, 0, 0},
		Value: []byte{1, 2, 2},
	}
	newque.Put(firstVal)
	newque.Put(secondVal)
	if !reflect.DeepEqual(newque.values[0], firstVal) {
		t.Error("first value in queue is not matching", newque.values[0])
	}
	if !reflect.DeepEqual(newque.values[1], secondVal) {
		t.Error("second value in queue is not matching")
	}
}
