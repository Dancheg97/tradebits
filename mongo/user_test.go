package mongo

import (
	"context"
	"testing"
	"time"
)

func TestUserCheck(t *testing.T) {
	time.Sleep(time.Millisecond * 200)
	Setup("mongodb://localhost:27017")
	UserCreate("12")
	exists := UserCheck("12")
	if !exists {
		t.Error("user does not exist")
	}
	userCollection.Drop(context.Background())
}

func TestUserCreate(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	Setup("mongodb://localhost:27017")
	err := UserCreate("123")
	if err != nil {
		t.Error(err)
	}
	userCollection.Drop(context.Background())
}

func TestUserGet(t *testing.T) {
	time.Sleep(time.Millisecond * 600)
	Setup("mongodb://localhost:27017")
	UserCreate("123")
	user, err := UserGet("123")
	if err != nil {
		t.Error(err)
	}
	if user.Pubkey != "123" {
		t.Error("wrong field")
	}
	userCollection.Drop(context.Background())
}
