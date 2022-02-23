package dgraph

import (
	"testing"
)

func TestOpenDB(t *testing.T) {
	db := openDB()
	if db == nil {
		t.Error("db should not be equal to nil")
	}
}
