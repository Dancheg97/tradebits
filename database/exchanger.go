package database

import "github.com/syndtr/goleveldb/leveldb"

type Exchanger struct {
	Name               string `json:"Name"`
	Image              []byte `json:"Image"`
	Pledge             uint64 `json:"Pledge"`
	MessageKey         []byte `json:"MessageKey"`
	GoodFeedBacksCount uint64 `json:"GoodFeedBacksCount"`
	BadFeedBacksCount  uint64 `json:"BadFeedBacksCount"`
	RequestsLink       []byte `json:"RequestsLink"`
}

var exchangerDB, _ = leveldb.OpenFile("database/exchangerData", nil)

func NewExchanger(adress []byte) {
	
}
