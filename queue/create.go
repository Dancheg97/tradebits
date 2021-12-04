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
		values: []KvPair{},
	}
	return &queue
}
