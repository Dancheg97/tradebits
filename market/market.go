package market

import (
	"bytes"
	"encoding/gob"
	"errors"
	"sync_tree/data"
	"sync_tree/lock"
)

type market struct {
	adress  []byte
	Name    string
	MesKey  []byte
	Descr   string
	Img     string
	OpCount uint64
	Buys    []Trade
	Sells   []Trade
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
		Buys:    []Trade{},
		Sells:   []Trade{},
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
Recursive function to add trades to existing market. Each new iteration
*/
func (m *market) OperateTrade(newTrade Trade) {
	m.OpCount = m.OpCount + 1
	if newTrade.IsSell {
		if len(m.Buys) == 0 {
			m.Sells = append(m.Sells, newTrade)
			return
		}
		trades, outputs := newTrade.operate(m.Buys[0])
		m.Buys = m.Buys[1:]
		m.outputs = append(m.outputs, outputs...)
		if len(trades) == 2 {
			m.addTrade(trades[0])
			m.addTrade(trades[1])
			return
		}
		if len(trades) == 1 {
			m.OperateTrade(newTrade)
		}
		return
	} else {
		if len(m.Sells) == 0 {
			m.Buys = append(m.Buys, newTrade)
			return
		}
		trades, outputs := newTrade.operate(m.Sells[0])
		m.Sells = m.Sells[1:]
		m.outputs = append(m.outputs, outputs...)
		if len(trades) == 2 {
			m.addTrade(trades[0])
			m.addTrade(trades[1])
			return
		}
		if len(trades) == 1 {
			m.OperateTrade(newTrade)
		}
		return
	}
}
