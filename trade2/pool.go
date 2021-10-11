package trade2

type tradePool struct {
	Buys  []trade
	Sells []trade
}

func CreatePool() *tradePool {
	return &tradePool{
		Buys:  []trade{},
		Sells: []trade{},
	}
}

func (tp *tradePool) insertSell(sell *trade) {
	currentRatio := float64(sell.Offer) / float64(sell.Recieve)
	for addIndex, checkSell := range tp.Sells {
		checkRatio := float64(checkSell.Offer) / float64(checkSell.Recieve)
		if currentRatio > checkRatio {
			tp.Sells = append(tp.Sells[:addIndex+1], tp.Sells[addIndex:]...)
			tp.Sells[addIndex] = *sell
			return
		}
	}
	tp.Sells = append(tp.Sells, *sell)
}

func (tp *tradePool) insertBuy(buy *trade) {
	currentRatio := float64(buy.Offer) / float64(buy.Recieve)
	for addIndex, checkBuy := range tp.Buys {
		checkRatio := float64(checkBuy.Offer) / float64(checkBuy.Recieve)
		if currentRatio > checkRatio {
			tp.Buys = append(tp.Buys[:addIndex+1], tp.Buys[addIndex:]...)
			tp.Buys[addIndex] = *buy
			return
		}
	}
	tp.Buys = append(tp.Buys, *buy)
}
