package user

import (
	"bytes"
	"encoding/gob"
	"sync_tree/__logs"
	"sync_tree/_data"
	"sync_tree/_lock"
)

type User interface {
	get(adress []byte) (*user, error)
	create(adress []byte, messageKey []byte, imageLink string)
	unlock()
	balance() uint64
	assetBalance(adress []byte) uint64
	messageKey() []byte
	imageLink() string
	changeBalance(balance uint64)
	changeAssetBalance(adress []byte, balance uint64)
	changeMessageKey(newPublicKeyBytes []byte)
	changeImageLink(string)
}

type user struct {
	mainBalance   uint64            `gob:"MainBalance"`
	messageKey    []byte            `gob:"MessageKey"`
	image         string            `gob:"Image"`
	assetBalances map[string]uint64 `gob:"AssetBalances"`
}

func (u user) get(adress []byte) (*user, error) {
	lock_err := _lock.Lock(adress)
	if lock_err != nil {
		return nil, __logs.Error("User is locked for another operation")
	}
	user := 
	userBytes := _data.Get(adress)
	
}

func 