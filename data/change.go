package data

// change existing value in database by key
func Change(key []byte, value []byte) {
	if Check(key) {
		base.Put(key, value, nil)
	}
	if reserveNodeConnected {
		obj := dataObject{
			key:   key,
			value: value,
		}
		DataQue = append(DataQue, obj)
	}
}
