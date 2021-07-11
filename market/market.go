package market

import (
	"bytes"
	"encoding/gob"
	"errors"
	"sync_tree/data"
	"sync_tree/lock"
	"time"
)

type market struct {
	adress  []byte
	Name    string
	MesKey  []byte
	Descr   string
	Img     string
	OpCount uint64
	// Buys    []Trade
	// Sells   []Trade
	Mes     map[string]string
	Arch    map[string]string
	outputs []output
}

type output struct {
	Adress    []byte
	MainOut   uint64
	MarketOut uint64
}

/*
Create new market by passed values. Checks wether market with passed adress
exists and creates new one.
*/
func Create(
	adress []byte,
	Name string,
	MesKey []byte,
	Descr string,
	Img string,
) error {
	if data.Check(adress) {
		return errors.New("possibly market already exists")
	}
	newMarket := market{
		adress:  adress,
		Name:    Name,
		Descr:   Descr,
		Img:     Img,
		MesKey:  MesKey,
		OpCount: 0,
		// Buys:    []Trade{},
		// Sells:   []Trade{},
		Mes:     make(map[string]string),
	}
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(newMarket)
	data.Put(adress, cache.Bytes())
	data.SearchAdd(Name, adress)
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
		time.Sleep(time.Millisecond * 89)
		return Get(adress)
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
func (a *market) Save() {
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

/*
Function to add message from some adress to concrete market
*/
func (m *market) PutMessage(userAdress []byte, mes string) {
	strAdr := string(userAdress)
	m.Mes[strAdr] = m.Mes[strAdr] + "|" + mes
}

/*
This function is made to get all new messages and to put all current messages
to archieve
*/
func (m *market) GetAllMessages() map[string]string {
	messages := m.Mes
	for sender, message := range m.Mes {
		m.Arch[sender] = m.Arch[sender] + "|" + message
	}
	m.Mes = make(map[string]string)
	return messages
}
