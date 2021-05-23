package user

import (
	"bytes"
	"encoding/gob"
	"sync_tree/__logs"
	"sync_tree/_data"
)

type User interface {
	Create(adress []byte, messageKey []byte, imageLink string) error
	Get(adress []byte) (*user, error)
	Unlock()
	Balance() uint64
	AssetBalance(adress []byte) uint64
	MessageKey() []byte
	ImageLink() string
	ChangeBalance(balance uint64)
	ChangeAssetBalance(adress []byte, balance uint64)
	ChangeMessageKey(newPublicKeyBytes []byte)
	ChangeImageLink(string)
}

type user struct {
	balance uint64            `gob:"b"`
	mesKey  []byte            `gob:"k"`
	img     string            `gob:"i"`
	assets  map[string]uint64 `gob:"a"`
}

func Create(adress []byte, mesKey []byte, img string) error {
	if _data.Check(adress) {
		return __logs.Error("create user existing key %v err", adress)
	}
	u := user{
		balance: 0,
		mesKey:  mesKey,
		img:     img,
		assets:  make(map[string]uint64),
	}
	userBytesBuffer := new(bytes.Buffer)
	gob.NewEncoder(userBytesBuffer).Encode(u)
	_data.Put(adress, userBytesBuffer.Bytes())
	__logs.Info("new user create success, adress: %v", adress)
	return nil
}
