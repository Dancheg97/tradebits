package asset

import (
	"bytes"
	"encoding/gob"
	"sync_tree/__logs"
	"sync_tree/_calc"
	"sync_tree/_data"
	"sync_tree/_lock"
)

type asset struct {
	adress   []byte
	Name     string
	ImgLink  string
	MesKey   []byte
	Likes    uint64
	Dislikes uint64
	Market   []byte
}

// Create new exchanger, in case there is already one with same adress
// or other technical troubles be logged
func Create(adress []byte, Name string, ImgLink string, MesKey []byte) error {
	if _data.Check(adress) {
		return __logs.Error("create asset by existing key: ", adress)
	}
	newAsset := asset{
		adress:   adress,
		Name:     Name,
		ImgLink:  ImgLink,
		MesKey:   MesKey,
		Likes:    0,
		Dislikes: 0,
		Market:   _calc.RandBytes(),
	}
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(newAsset)
	_data.Put(adress, cache.Bytes())
	__logs.Info("new user create success, adress: ", adress)
	return nil
}

// Get asset to make some changes with it's contents, that func is blocking
// so use it only when making some changes to asset data
func Get(adress []byte) *asset {
	lockErr := _lock.Lock(adress)
	if lockErr != nil {
		__logs.Error("unable to get asset, locked: ", adress)
		return nil
	}
	u := asset{adress: adress}
	userBytes := _data.Get(adress)
	cache := bytes.NewBuffer(userBytes)
	gob.NewDecoder(cache).Decode(&u)
	return &u
}
