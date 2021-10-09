package data

// change existing value in database by key
func Change(key []byte, value []byte) {
	if Check(key) {
		base.Put(key, value, nil)
	}
}
