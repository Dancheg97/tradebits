package market

import (
	"bytes"
	"encoding/gob"
	"sync_tree/data"
	"sync_tree/lock"
	"sync_tree/logs"
)

type market struct {
	adress   []byte
	Name     string
	ImgLink  string
	MesKey   []byte
	Likes    uint64
	Dislikes uint64
	Buys     []Buy
	Sells    []Sell
}

/*
Create new market by passed values. Checks wether market with passed adress
exists and creates new one.
*/
func Create(adress []byte, Name string, ImgLink string, MesKey []byte) error {
	if data.Check(adress) {
		return logs.Error("create market by existing key: ", adress)
	}
	newMarket := market{
		adress:   adress,
		Name:     Name,
		ImgLink:  ImgLink,
		MesKey:   MesKey,
		Likes:    0,
		Dislikes: 0,
		Buys:     []Buy{},
		Sells:    []Sell{},
	}
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(newMarket)
	data.Put(adress, cache.Bytes())
	logs.Info("new user create success, adress: ", adress)
	return nil
}

/*
This function is blocking, it gives an instance of market, so that the
values of that market can be modified. To save changes to DB call Save().

Only one instance of market can be called at same time.

This function should be used only in case those values are modified:
 - Name
 - ImgLink
 - MesKey
 - Likes
 - DisLikes
*/
func Get(adress []byte) *market {
	lockErr := lock.Lock(adress)
	if lockErr != nil {
		logs.Error("unable to get market, locked: ", adress)
		return nil
	}
	a := market{adress: adress}
	marketBytes := data.Get(adress)
	cache := bytes.NewBuffer(marketBytes)
	gob.NewDecoder(cache).Decode(&a)
	return &a
}

/*
This function is saving changes to the market in database and removes ability
to make a double save by removing adress from class struct.
*/
func (a market) Save() {
	saveAdress := a.adress
	a.adress = nil
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(a)
	data.Change(saveAdress, cache.Bytes())
	lock.Unlock(saveAdress)
}

/*
Non blocking function to look for market contents, it's impossible to save
instance of that market to database.
*/
func Look(adress []byte) *market {
	currMarket := market{}
	marketBytes := data.Get(adress)
	marketCache := bytes.NewBuffer(marketBytes)
	gob.NewDecoder(marketCache).Decode(&currMarket)
	return &currMarket
}
