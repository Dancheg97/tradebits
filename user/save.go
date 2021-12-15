package user

import (
	"bytes"
	"encoding/gob"
	"orb/database"
	"orb/lock"
)

/*
This function is used to save user, after that function his state is
gonna be fixed in databasebase, adress will be unlocked, and adress will be
set to nil, so other changes won't be saved (user will have to be recreated)
*/
func (u *user) Save() {
	saveAdress := u.adress
	u.adress = nil
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(u)
	database.Change(saveAdress, cache.Bytes())
	lock.Unlock(saveAdress)
}
