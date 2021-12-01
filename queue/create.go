package queue

import "sync"

type queue struct {
	mu     sync.Mutex
	values []KvPair
}

type KvPair struct {
	Key   []byte
	Value []byte
}

func Create() *queue {
	queue := queue{
		values: make([]KvPair, 10000),
	}
	return &queue
}

func (qu *queue) Take() KvPair {
	qu.mu.Lock()
	pair := qu.values[0]
	qu.values = qu.values[1:]
	qu.mu.Unlock()
	return pair
}
