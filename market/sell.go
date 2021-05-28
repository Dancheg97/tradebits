package market

import "math"

type Buy struct {
	Adress        []byte
	OfferMain     uint64
	RecieveMarket uint64
}

type Sell struct {
	Adress      []byte
	OfferMarket uint64
	RecieveMain uint64
}

func (sell Sell) checkMatch(buy Buy) bool {
	sellRatio := float64(buy.RecieveMarket / buy.OfferMain)
	buyRatio := float64(sell.OfferMarket / sell.RecieveMain)
	return sellRatio < buyRatio
}

func (sell Sell) selfClose(buy Buy) bool {
	return sell.RecieveMain < buy.OfferMain
}
