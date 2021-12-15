package market

import (
	"bytes"
	"encoding/gob"
	"orb/database"
	"orb/lock"
)

// This function is saving changes to the market in databasebase and removes ability
// to make a double save by removing adress from class struct.
func (m *market) Save() {
	saveAdress := m.adress
	m.adress = nil
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(m)
	database.Change(saveAdress, cache.Bytes())
	lock.Unlock(saveAdress)
}
