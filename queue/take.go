package queue

func (qu *queue) Take() KvPair {
	qu.mu.Lock()
	pair := qu.values[0]
	qu.values = qu.values[1:]
	qu.mu.Unlock()
	return pair
}
