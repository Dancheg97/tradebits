package market

import (
	"bytes"
	"encoding/gob"
	"orb/database"
)

// Non blocking function to look for market contents, it's impossible to save
// instance of that market to databasebase.
func Look(adress []byte) *market {
	currMarket := market{}
	marketBytes := database.Get(adress)
	marketCache := bytes.NewBuffer(marketBytes)
	gob.NewDecoder(marketCache).Decode(&currMarket)
	return &currMarket
}
