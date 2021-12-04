package market

import (
	"bytes"
	"encoding/gob"
	"sync_tree/data"
	"sync_tree/lock"
)

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
