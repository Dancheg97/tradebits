package database

import (
	"bytes"
	"encoding/gob"
	"sync_tree/__logs"
	"sync_tree/_calc"
	"sync_tree/_data"
)

type asset struct {
	adress   []byte
	Name     string
	ImgLink  []byte
	MesKey   []byte
	Likes    []byte
	Dislikes []byte
	Trades   []byte
}

// Create new exchanger, in case there is already one with same adress
// or other technical troubles be logged
func Create(adress []byte, Name string, ImgLink []byte, MesKey []byte) error {
	if _data.Check(adress) {
		return __logs.Error("create asset by existing key ", adress)
	}
	newAsset := asset{
		adress:   adress,
		Name:     Name,
		ImgLink:  ImgLink,
		MesKey:   MesKey,
		Likes:    _calc.RandBytes(),
		Dislikes: _calc.RandBytes(),
		Trades:   _calc.RandBytes(),
	}
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(newAsset)
	_data.Put(adress, cache.Bytes())
	__logs.Info("new user create success, adress: %v", adress)
	return nil
}
