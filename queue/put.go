package queue

func (qu *keyValueQueue) Put(kvpair KvPair) {
	qu.mu.Lock()
	qu.values = append(qu.values, kvpair)
	qu.mu.Unlock()
}

func (qu *requestQueue) Put(request Request) {
	qu.mu.Lock()
	qu.values = append(qu.values, request)
	qu.mu.Unlock()

}
