package queue

import "testing"

func TestCreateQueue(t *testing.T) {
	newque := CreateKV()
	if newque == nil {
		t.Error("queue should not be returned as nil")
		return
	}
	if newque.values == nil {
		t.Error("queue values have not been initilized")
	}
	if len(newque.values) != 0 {
		t.Error("queue values length incorrect")
	}
}
