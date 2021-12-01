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
		mu:     sync.Mutex{},
		values: make([]KvPair, 1000),
	}
	return &queue
}

func (qu *queue) Put(kvpair KvPair) {
	qu.mu.Lock()
	qu.values = append(qu.values, kvpair)
	qu.mu.Unlock()
}

func (qu *queue) Take() KvPair {
	qu.mu.Lock()
	pair := qu.values[0]
	qu.values = qu.values[1:]
	qu.mu.Unlock()
	return pair
}
