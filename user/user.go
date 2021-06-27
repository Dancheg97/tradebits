package user

import (
	"bytes"
	"encoding/gob"
	"errors"
	"sync_tree/data"
	"sync_tree/lock"
)

type user struct {
	adress     []byte
	Balance    uint64
	MesKey     []byte
	PublicName string
	Markets    map[string]uint64
	Mes        map[string]string
	Arch       map[string]string
}

/*
Create new user, in case there is already user with same adress
the error will be logged
*/
func Create(adress []byte, MesKey []byte, PublicName string) error {
	if data.Check(adress) {
		return errors.New("possibly user already exists")
	}
	u := user{
		Balance:    0,
		MesKey:     MesKey,
		PublicName: PublicName,
		Markets:    make(map[string]uint64),
		Mes:        make(map[string]string),
		Arch:       map[string]string{},
	}
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(u)
	data.Put(adress, cache.Bytes())
	return nil
}

/*
Get existing user from database, by getting user his ID is gonna be
locked, so another of that user are not gonna appear
*/
func Get(adress []byte) *user {
	lockErr := lock.Lock(adress)
	if lockErr != nil {
		return nil
	}
	u := user{adress: adress}
	userBytes := data.Get(adress)
	cache := bytes.NewBuffer(userBytes)
	gob.NewDecoder(cache).Decode(&u)
	return &u
}

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

// Get user balance for some specific market
func (u *user) MarketBalance(market []byte) uint64 {
	return u.Markets[string(market)]
}

// Change user balance for some specific market
func (u *user) ChangeMarketBalance(market []byte, balance uint64) {
	u.Markets[string(market)] = balance
}

/*
Function to add message from some adress to concrete user
*/
func (u *user) PutMessage(userAdress []byte, mes string) {
	strAdr := string(userAdress)
	u.Mes[strAdr] = u.Mes[strAdr] + "|" + mes
}

/*
This function is made to get all new messages and to put all current messages
to archieve
*/
func (u *user) GetAllMessages() map[string]string {
	messages := u.Mes
	for sender, message := range u.Mes {
		u.Arch[sender] = u.Arch[sender] + "|" + message
	}
	u.Mes = make(map[string]string)
	return messages
}
