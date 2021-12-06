package queue

import (
	"reflect"
	"testing"
)

func TestTake(t *testing.T) {
	que := Create()
	first := KvPair{
		Key:   []byte{0},
		Value: []byte{0},
	}
	second := KvPair{
		Key:   []byte{1},
		Value: []byte{1},
	}
	third := KvPair{
		Key:   []byte{2},
		Value: []byte{2},
	}
	que.Put(first)
	que.Put(second)
	que.Put(third)
	firstTaken := que.Take()
	secondTaken := que.Take()
	thirdTaken := que.Take()
	if !reflect.DeepEqual(first, firstTaken) {
		t.Error("first taken message is incorrect", firstTaken)
	}
	if !reflect.DeepEqual(second, secondTaken) {
		t.Error("first taken message is incorrect", secondTaken)
	}
	if !reflect.DeepEqual(third, thirdTaken) {
		t.Error("first taken message is incorrect", thirdTaken)
	}
	if len(que.values) != 0 {
		t.Error("something is left in que")
	}
}

func TestTryToTakeFromEmptyQueue(t *testing.T) {
	que := Create()
	que.Take()
	
}