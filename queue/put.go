package queue

func (qu *keyValueQueue) Put(kvpair KvPair) {
	qu.mu.Lock()
	qu.values = append(qu.values, kvpair)
	qu.mu.Unlock()
}
