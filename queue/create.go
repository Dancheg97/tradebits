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
	Type           string
	PublicKey      []byte
	MessageKey     []byte
	PublicName     string
	SendAmount     uint64
	RecieverAdress []byte
	Adress         []byte
	Message        string
	Recieve        []byte
	Offer          uint64
	MarketAdress   []byte
	Sign           []byte
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
