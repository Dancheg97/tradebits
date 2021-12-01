package queue

import "testing"

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
	
}
