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

func CreateQueue() *queue {
	queue := queue{
		mu:     sync.Mutex{},
		values: make([]KvPair, 1000),
	}
	return &queue
}
