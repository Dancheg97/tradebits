package user

import (
	"bytes"
	"encoding/gob"
	"sync_tree/__logs"
	"sync_tree/_data"
	"sync_tree/_lock"
)

type user struct {
	adress  []byte
	Balance uint64
	MesKey  []byte
	ImgLink string
	Assets  map[string]uint64
}

/*
Create new user, in case there is already user with same adress
the error will be logged
*/
func Create(adress []byte, MesKey []byte, ImgLink string) error {
	if _data.Check(adress) {
		return __logs.Error("create user by existing key ", adress)
	}
	u := user{
		Balance: 0,
		MesKey:  MesKey,
		ImgLink: ImgLink,
		Assets:  make(map[string]uint64),
	}
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(u)
	_data.Put(adress, cache.Bytes())
	__logs.Info("new user create success, adress: ", adress)
	return nil
}

/*
Get existing user from database, by getting user his ID is gonna be
locked, so another of that user are not gonna appear
*/
func Get(adress []byte) *user {
	lockErr := _lock.Lock(adress)
	if lockErr != nil {
		__logs.Error("unable to get user (locked): ", adress)
		return nil
	}
	u := user{adress: adress}
	userBytes := _data.Get(adress)
	cache := bytes.NewBuffer(userBytes)
	gob.NewDecoder(cache).Decode(&u)
	return &u
}

/*
This function is used to save user, after that function his state is
gonna be fixed in database, adress will be unlocked, and adress will be
set to nil, so other changes won't be saved (user will have to be recreated)
*/
func (u user) Save() {
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(u)
	unlock_adress := u.adress
	_data.Change(u.adress, cache.Bytes())
	u.adress = nil
	_lock.Unlock(unlock_adress)
}

// Get user balance for some specific asset
func (u user) AssetBalance(asset []byte) uint64 {
	return u.Assets[string(asset)]
}

// Change user balance for some specific asset
func (u user) ChangeAssetBalance(asset []byte, balance uint64) {
	u.Assets[string(asset)] = balance
}
