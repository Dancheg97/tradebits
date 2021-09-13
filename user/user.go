package user

import (
	"bytes"
	"encoding/gob"
	"errors"
	"sync_tree/data"
	"sync_tree/lock"
	"sync_tree/trade"
	"time"
)

type user struct {
	adress     []byte
	Balance    uint64
	MesKey     []byte
	PublicName string
	Balances   map[string]uint64
	Messages   map[string][]string
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
		Balances:   make(map[string]uint64),
		Messages:   make(map[string][]string),
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
		time.Sleep(time.Millisecond * 89)
		return Get(adress)
	}
	u := user{adress: adress}
	userExists := data.Check(adress)
	if userExists == false {
		return nil
	}
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

/*
Function to add message from some adress to concrete user
*/
func (u *user) PutUserMessage(adress []byte, mes string) {
	strAdress := string(adress)
	if u.Messages[strAdress] == nil {
		u.Messages[strAdress] = []string{"u" + mes}
		return
	}
	u.Messages[strAdress] = append(u.Messages[strAdress], mes)
}

func (u *user) PutMarketMessage(adress []byte, mes string) {
	strAdress := string(adress)
	if u.Messages[strAdress] == nil {
		u.Messages[strAdress] = []string{"m" + mes}
		return
	}
	u.Messages[strAdress] = append(u.Messages[strAdress], mes)
}

/*
This function is made to get all new messages and to put all current messages
to archieve
*/
func (u *user) GetMessages(adress []byte) []string {
	return u.Messages[string(adress)]
}

// This function is bounding specific sell function to user, if its not
// possible, returns false, if operated successully, returns true
func (u *user) AttachSell(sell *trade.Sell, adress []byte) bool {
	if u.adress == nil {
		return false
	}
	if sell.Adress != nil {
		return false
	}
	if sell.Offer == 0 || sell.Recieve == 0 {
		return false
	}
	if val, ok := u.Balances[string(adress)]; ok {
		if val >= sell.Offer {
			u.Balances[string(adress)] = val - sell.Offer
			sell.Adress = u.adress
			return true
		}
	}
	return false
}

// this function is bounding specific buy to user, if its not possible returns
// false, if operated successfully returns true
func (u *user) AttachBuy(buy *trade.Buy) bool {
	if u.adress == nil {
		return false
	}
	if buy.Adress != nil {
		return false
	}
	if buy.Offer == 0 || buy.Recieve == 0 {
		return false
	}
	if u.Balance < buy.Offer {
		return false
	}
	u.Balance = u.Balance - buy.Offer
	buy.Adress = u.adress
	return true
}
