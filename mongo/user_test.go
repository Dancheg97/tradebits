package mongo

import "testing"

func TestNewUserSuccess(t *testing.T) {
	Setup("mongodb://localhost:27017")
	user := User{
		PubKey:   "123",
		Balance:  0,
		Messages: []string{},
	}
	err := NewUser(user)
	if err != nil {
		t.Error(err)
	}
}
