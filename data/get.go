package data

// get value by key from database
func Get(key []byte) []byte {
	output, getErr := base.Get(key, nil)
	if getErr != nil {
		return nil
	}
	return output
}
