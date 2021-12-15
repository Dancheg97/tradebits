package user

import (
	"bytes"
	"encoding/gob"
	"orb/data"
)

/*
Non blocking function to look for user contents, it's impossible to save
instance of that user to database.
*/
func Look(adress []byte) *user {
	u := user{}
	userBytes := data.Get(adress)
	cache := bytes.NewBuffer(userBytes)
	gob.NewDecoder(cache).Decode(&u)
	return &u
}
