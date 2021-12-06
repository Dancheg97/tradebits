package queue

import "sync"

type keyValueQueue struct {
	mu     sync.Mutex
	values []KvPair
}

type requestQueue struct {
	mu     sync.Mutex
	values []Request
}

type Request struct {
}

type KvPair struct {
	Key   []byte
	Value []byte
}

func CreateKV() *keyValueQueue {
	queue := keyValueQueue{
		values: []KvPair{},
	}
	return &queue
}

func CreateREQ() *requestQueue {
	requestQueue := requestQueue{
		values: []Request{},
	}
	return &requestQueue
}
