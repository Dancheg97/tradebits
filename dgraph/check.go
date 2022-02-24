package data

// check if value exists in database
func Check(key []byte) bool {
	valueExists, _ := base.Has(key, nil)
	return valueExists
}
