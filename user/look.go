package user

import (
	"bytes"
	"encoding/gob"
	"orb/database"
)

/*
Non blocking function to look for user contents, it's impossible to save
instance of that user to databasebase.
*/
func Look(adress []byte) *user {
	u := user{}
	userBytes := database.Get(adress)
	cache := bytes.NewBuffer(userBytes)
	gob.NewDecoder(cache).Decode(&u)
	return &u
}
