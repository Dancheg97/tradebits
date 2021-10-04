package market

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"reflect"
	"sync_tree/data"
	"sync_tree/lock"
	"sync_tree/search"
	"sync_tree/trade"
	"sync_tree/user"
)

type market struct {
	adress    []byte
	Name      string
	MesKey    []byte
	Descr     string
	Img       string
	OpCount   uint64
	Pool      trade.TradePool
	InputFee  uint64
	OutputFee uint64
	WorkTime  string
	Delimiter uint64
	Users     [][]byte
}

// Create new market by passed values. Checks wether market with passed adress
// exists and creates new one. Here is field description:
//
// - Adress: represens hash of markets public key
//
// - Name: market name visible for users (min 10, max 30)
//
// - MesKey: message key that is gonna be used to check
//
// - Descr: market description visible for users (min 160, max 480)
//
// - Img: url link to the image, bcs market dont store images
//
// - InputFee: fee value, each round number representing 0.01% (min 0, max 500)
//
// - OutputFee: fee value, each round number representing 0.01% (min 0, max 500)
//
// - WorkTime: representing when market is working with messages (min15, max 45)
//
// - Delimiter: value that is representing decimal places of its value (max 10)
func Create(
	adress []byte,
	name string,
	mesKey []byte,
	descr string,
	imgLink string,
	inputFee uint64,
	outputFee uint64,
	workTime string,
	delimiter uint64,
) error {
	if len(adress) != 64 {
		return errors.New("bad adress length")
	}
	if len(name) < 10 || len(name) > 30 {
		return errors.New("bad name length")
	}
	if len(mesKey) < 240 || len(mesKey) > 320 {
		fmt.Println(len(mesKey))
		return errors.New("invalid message key length")
	}
	if len(descr) < 160 || len(descr) > 760 {
		return errors.New("bad description length")
	}
	if inputFee > 500 || outputFee > 500 {
		return errors.New("fee too big")
	}
	if len(workTime) < 10 || len(workTime) > 40 {
		return errors.New("work time is bad")
	}
	if delimiter > 10 {
		return errors.New("delimiter length is too long")
	}
	if data.Check(adress) {
		return errors.New("possibly market already exists")
	}
	if data.Check([]byte(name)) {
		return errors.New("market with that name exists")
	}
	data.Put([]byte(name), []byte{})
	pool := trade.TradePool{
		Buys:    []trade.Buy{},
		Sells:   []trade.Sell{},
		Outputs: []trade.Output{},
	}
	newMarket := market{
		adress:    adress,
		Name:      name,
		Descr:     descr,
		Img:       imgLink,
		MesKey:    mesKey,
		OpCount:   0,
		Pool:      pool,
		InputFee:  inputFee,
		OutputFee: outputFee,
		WorkTime:  workTime,
		Delimiter: delimiter,
		Users:     [][]byte{},
	}
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(newMarket)
	data.Put(adress, cache.Bytes())
	search.Add(name, adress)
	return nil
}

// This function is blocking, it gives an instance of market, so that the
// values of that market can be modified. To save changes made in market call
// Save() method of returned instance.
func Get(adress []byte) *market {
	if !data.Check(adress) {
		return nil
	}
	lock.Lock(adress)
	m := market{adress: adress}
	marketBytes := data.Get(adress)
	cache := bytes.NewBuffer(marketBytes)
	gob.NewDecoder(cache).Decode(&m)
	return &m
}

// This function is saving changes to the market in database and removes ability
// to make a double save by removing adress from class struct.
func (m *market) Save() {
	saveAdress := m.adress
	m.adress = nil
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(m)
	data.Change(saveAdress, cache.Bytes())
	lock.Unlock(saveAdress)
}

// Non blocking function to look for market contents, it's impossible to save
// instance of that market to database.
func Look(adress []byte) *market {
	currMarket := market{}
	marketBytes := data.Get(adress)
	marketCache := bytes.NewBuffer(marketBytes)
	gob.NewDecoder(marketCache).Decode(&currMarket)
	return &currMarket
}

// This function is operating output for some trade and market adress
func operateOutput(t trade.Output, adress []byte) {
	u := user.Get(t.Adress)
	u.Balance = u.Balance + t.Main
	u.Balances[string(adress)] = u.Balances[string(adress)] + t.Market
	u.Save()
}

// attaches buy trade to market, you can't attach trade twice
func (m *market) AttachBuy(b *trade.Buy) bool {
	if m.adress == nil {
		return false
	}
	if b.Adress == nil {
		return false
	}
	m.Pool.OperateBuy(*b)
	for _, output := range m.Pool.Outputs {
		go operateOutput(output, m.adress)
	}
	m.Pool.Outputs = []trade.Output{}
	b.Adress = nil
	b.Offer = 0
	return true
}

// attaches sell trade to market, you can't attach trade twice
func (m *market) AttachSell(s *trade.Sell) bool {
	if m.adress == nil {
		return false
	}
	if s.Adress == nil {
		return false
	}
	m.Pool.OperateSell(*s)
	for _, output := range m.Pool.Outputs {
		go operateOutput(output, m.adress)
	}
	m.Pool.Outputs = []trade.Output{}
	s.Adress = nil
	s.Offer = 0
	return true
}

// making a check, wether some user has trades on that market
func (m *market) HasTrades(adress []byte) bool {
	for _, trade := range m.Pool.Buys {
		if reflect.DeepEqual(trade.Adress, adress) {
			return true
		}
	}
	for _, trade := range m.Pool.Sells {
		if reflect.DeepEqual(trade.Adress, adress) {
			return true
		}
	}
	return false
}

func cancelBuy(userAdress []byte, offer uint64) {
	usr := user.Get(userAdress)
	usr.Balance = usr.Balance + offer
	usr.Save()
}

func cancelSell(userAdress []byte, marketAdress []byte, offer uint64) {
	usr := user.Get(userAdress)
	mktAdress := string(marketAdress)
	usr.Balances[mktAdress] = usr.Balances[mktAdress] + offer
	usr.Save()
}

// this function cancelles trades
func (m *market) CancelTrades(adress []byte) {
	for idx, trade := range m.Pool.Buys {
		if reflect.DeepEqual(trade.Adress, adress) {
			go cancelBuy(adress, trade.Offer)
			m.Pool.Buys = append(m.Pool.Buys[:idx], m.Pool.Buys[idx+1:]...)
		}
	}
	for idx, trade := range m.Pool.Sells {
		if reflect.DeepEqual(trade.Adress, adress) {
			go cancelSell(adress, m.adress, trade.Offer)
			m.Pool.Sells = append(m.Pool.Sells[:idx], m.Pool.Sells[idx+1:]...)
		}
	}
}
