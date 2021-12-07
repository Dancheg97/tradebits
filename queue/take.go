package queue

func (que *keyValueQueue) Take() KvPair {
	que.mu.Lock()
	if len(que.values) == 0 {
		return KvPair{}
	}
	pair := que.values[0]
	que.values = que.values[1:]
	que.mu.Unlock()
	return pair
}

func (que *requestQueue) Take() Request {
	que.mu.Lock()
	if len(que.values) == 0 {
		return Request{}
	}
	request := que.values[0]
	que.values = que.values[1:]
	que.mu.Unlock()
	return request
}
