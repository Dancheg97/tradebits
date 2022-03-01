package mongo

import "testing"

func TestNewUserSuccess(t *testing.T) {
	Setup("mongodb://localhost:27017")
	err := UserCreate("user")
	if err != nil {
		t.Error(err)
	}
}
