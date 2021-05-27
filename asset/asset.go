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

// request to buy some asset
type Buy struct {
	Adress       []byte
	OfferMain    uint64
	RecieveAsset uint64
}

// request to sell some asset
type Sell struct {
	Adress      []byte
	OfferAsset  uint64
	RecieveMain uint64
}

// struct containing info about outputs
type Output struct {
	Adress   []byte
	MainOut  uint64
	AssetOut uint64
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

// checker: if 2 trades matches
func CheckMatch(sell Sell, buy Buy) bool {
	return float64(buy.RecieveAsset/buy.OfferMain) < float64(sell.OfferAsset/sell.RecieveMain)
}

// checker: if trade closed on buy side
func IfCloseBuyer(sell Sell, buy Buy) bool {
	return sell.RecieveMain > buy.OfferMain
}

/*
This function is taking sell and buy transaction, and closing sell transaction,
giving out the rest of buy transaction and. Outputs:
 - new Buy request for buyer
 - output for seller
 - output for buyer
*/
func CloseSeller(sell Sell, buy Buy) (Buy, Output, Output) {
	goesToSeller := uint64(buy.RecieveAsset / buy.OfferMain * sell.RecieveMain)
	sellerOutput := Output{
		AssetOut: goesToSeller,
		MainOut:  sell.RecieveMain,
	}
	buyerOutput := Output{
		AssetOut: sell.OfferAsset,
	}
	buy.OfferMain = sell.RecieveMain - buy.OfferMain
	buy.RecieveAsset = buy.RecieveAsset - goesToSeller
	return buy, sellerOutput, buyerOutput
}

