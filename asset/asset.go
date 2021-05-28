package asset

import (
	"bytes"
	"encoding/gob"
	"sync_tree/__logs"
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
	Buys     []Buy
	Sells    []Sell
}


/*
Create new asset by passed values. Checks wether asset with passed adress
exists and creates new one.
*/
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
		Buys:     []Buy{},
		Sells:    []Sell{},
	}
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(newAsset)
	_data.Put(adress, cache.Bytes())
	__logs.Info("new user create success, adress: ", adress)
	return nil
}

/*
This function is blocking, it gives an instance of asset, so that the
values of that asset can be modified. To save changes to DB call Save().

Only one instance of asset can be called at same time.

This function should be used only in case those values are modified:
 - Name
 - ImgLink
 - MesKey
 - Likes
 - DisLikes
*/
func Get(adress []byte) *asset {
	lockErr := _lock.Lock(adress)
	if lockErr != nil {
		__logs.Error("unable to get asset, locked: ", adress)
		return nil
	}
	a := asset{adress: adress}
	assetBytes := _data.Get(adress)
	cache := bytes.NewBuffer(assetBytes)
	gob.NewDecoder(cache).Decode(&a)
	return &a
}

/*
This function is saving changes to the asset in database and removes ability
to make a double save by removing adress from class struct.
*/
func (a asset) Save() {
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(a)
	unlock_adress := a.adress
	_data.Change(a.adress, cache.Bytes())
	a.adress = nil
	_lock.Unlock(unlock_adress)
}

/*
Non blocking function to look for asset contents, it's impossible to save
instance of that asset to database.
*/
func Look(adress []byte) *asset {
	currAsset := asset{}
	assetBytes := _data.Get(adress)
	assetCache := bytes.NewBuffer(assetBytes)
	gob.NewDecoder(assetCache).Decode(&currAsset)
	return &currAsset
}
