package market

import (
	"bytes"
	"encoding/gob"
	"sync_tree/data"
)

// Non blocking function to look for market contents, it's impossible to save
// instance of that market to database.
func Look(adress []byte) *market {
	currMarket := market{}
	marketBytes := data.Get(adress)
	marketCache := bytes.NewBuffer(marketBytes)
	gob.NewDecoder(marketCache).Decode(&currMarket)
	return &currMarket
}
