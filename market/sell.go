package market

type Trade struct {
	Adress        []byte
	OfferMarket   uint64
	OfferMain     uint64
	RecieveMain   uint64
	RecieveMarket uint64
}

func (sell Trade) checkMatch(buy Trade) bool {
	sellRatio := float64(buy.RecieveMarket / buy.OfferMain)
	buyRatio := float64(sell.OfferMarket / sell.RecieveMain)
	return sellRatio < buyRatio
}

func (sell Trade) selfClose(buy Trade) bool {
	return sell.RecieveMain < buy.OfferMain
}

func (sell *Trade) closeSell(buy Trade) {

}

func (sell Trade) closeBuy(buy Trade) {

}
