package market

import (
	"bytes"
	"encoding/gob"
	"orb/database"
	"orb/lock"
)

// This function is blocking, it gives an instance of market, so that the
// values of that market can be modified. To save changes made in market call
// Save() method of returned instance.
func Get(adress []byte) *market {
	if !database.Check(adress) {
		return nil
	}
	lock.Lock(adress)
	m := market{adress: adress}
	marketBytes := database.Get(adress)
	cache := bytes.NewBuffer(marketBytes)
	gob.NewDecoder(cache).Decode(&m)
	return &m
}
