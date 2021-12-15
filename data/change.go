package database

// change existing value in databasebase by key
func Change(key []byte, value []byte) {
	if Check(key) {
		base.Put(key, value, nil)
	}
	if reserveNodeConnected {
		obj := databaseObject{
			key:   key,
			value: value,
		}
		databaseQue = append(databaseQue, obj)
	}
}
