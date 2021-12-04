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
	if !reflect.DeepEqual(que.values[0], firstTaken) {
		t.Error("first message is not in queue/not in right position")
	}
	if !reflect.DeepEqual(que.values[1], secondTaken) {
		t.Error("second message is not in queue/not in right position")
	}
	if !reflect.DeepEqual(que.values[2], thirdTaken) {
		t.Error("third message is not in queue/not in right position")
	}
	if len(que.values) != 0 {
		t.Error("something is left in que")
	}
}
