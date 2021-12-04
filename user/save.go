package user

import (
	"bytes"
	"encoding/gob"
	"sync_tree/data"
	"sync_tree/lock"
)

/*
This function is used to save user, after that function his state is
gonna be fixed in database, adress will be unlocked, and adress will be
set to nil, so other changes won't be saved (user will have to be recreated)
*/
func (u *user) Save() {
	saveAdress := u.adress
	u.adress = nil
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(u)
	data.Change(saveAdress, cache.Bytes())
	lock.Unlock(saveAdress)
}
