package database

// check if value exists in databasebase
func Check(key []byte) bool {
	valueExists, _ := base.Has(key, nil)
	return valueExists
}
