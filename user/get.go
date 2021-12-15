package user

import (
	"bytes"
	"encoding/gob"
	"orb/database"
	"orb/lock"
)

/*
Get existing user from databasebase, by getting user his ID is gonna be
locked, so another of that user are not gonna appear
*/
func Get(adress []byte) *user {
	if len(adress) != 64 {
		return nil
	}
	if !database.Check(adress) {
		return nil
	}
	lock.Lock(adress)
	u := user{adress: adress}
	userBytes := database.Get(adress)
	cache := bytes.NewBuffer(userBytes)
	gob.NewDecoder(cache).Decode(&u)
	return &u
}
