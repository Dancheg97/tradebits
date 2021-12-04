package user

import (
	"bytes"
	"encoding/gob"
	"sync_tree/data"
	"sync_tree/lock"
)

/*
Get existing user from database, by getting user his ID is gonna be
locked, so another of that user are not gonna appear
*/
func Get(adress []byte) *user {
	if len(adress) != 64 {
		return nil
	}
	if !data.Check(adress) {
		return nil
	}
	lock.Lock(adress)
	u := user{adress: adress}
	userBytes := data.Get(adress)
	cache := bytes.NewBuffer(userBytes)
	gob.NewDecoder(cache).Decode(&u)
	return &u
}
