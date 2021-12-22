package chat

import (
	"orb/data"
	"testing"
)

func TestCreateChatMarketError(t *testing.T) {
	err := createChat([]byte{0, 1, 2, 3}, []byte{0, 1, 2, 3})
	if err == nil {
		t.Error("this test should have failed with market error")
	}
}

func TestCreateChatUserError(t *testing.T) {
	data.Put([]byte{0, 1, 2}, []byte{})
	err := createChat([]byte{0, 1, 2}, []byte{0, 1, 1})
	if err == nil {
		t.Error("this test should not be passed with user error")
	}
	data.TestRM([]byte{0, 1, 2})
}

func TestCreateChatExistsError(t *testing.T) {
	data.Put([]byte{0, 4, 2}, []byte{})
	data.Put([]byte{0, 4, 3}, []byte{})
	createChat([]byte{0, 4, 2}, []byte{0, 4, 3})
	err := createChat([]byte{0, 4, 2}, []byte{0, 4, 3})
	if err == nil {
		t.Error("second chat should not be created due to an error")
	}
	data.TestRM([]byte{0, 4, 2})
	data.TestRM([]byte{0, 4, 3})
}

func TestCreateChat(t *testing.T) {
	data.Put([]byte{0, 8, 2}, []byte{})
	data.Put([]byte{0, 8, 3}, []byte{})
	err := createChat([]byte{0, 8, 2}, []byte{0, 8, 3})
	if err != nil {
		t.Error("this creation should not throw an error")
	}
	data.TestRM([]byte{0, 8, 2})
	data.TestRM([]byte{0, 8, 3})
}
