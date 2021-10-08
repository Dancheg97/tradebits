package market

import (
	"bytes"
	"encoding/gob"
	"sync_tree/data"
	"sync_tree/lock"
)

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
