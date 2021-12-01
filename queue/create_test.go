package queue

import "testing"

func TestCreateQueue(t *testing.T) {
	newque := Create()
	if newque == nil {
		t.Error("queue should not be returned as nil")
	}
	if newque.values == nil {
		t.Error("queue values have not been initilized")
	}
	if len(newque.values) != 10000 {
		t.Error("queue is created with a wrong length")
	}
}
