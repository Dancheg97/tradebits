package trade2

type tradePool struct {
	Buys          []trade
	MainOutputs   []output
	Sells         []trade
	MarketOutputs []output
}

func CreatePool() *tradePool {
	return &tradePool{
		Buys:  []trade{},
		Sells: []trade{},
	}
}

func (pool *tradePool) insertSell(sell *trade) {
	currentRatio := float64(sell.Offer) / float64(sell.Recieve)
	for addIndex, checkSell := range pool.Sells {
		checkRatio := float64(checkSell.Offer) / float64(checkSell.Recieve)
		if currentRatio > checkRatio {
			pool.Sells = append(
				pool.Sells[:addIndex+1],
				pool.Sells[addIndex:]...,
			)
			pool.Sells[addIndex] = *sell
			return
		}
	}
	pool.Sells = append(pool.Sells, *sell)
}

func (pool *tradePool) ejectFirstBuy() *trade {
	buy := pool.Buys[0]
	pool.Buys = pool.Buys[1:]
	return &buy
}

func (pool *tradePool) OperateSell(sell *trade) {
	if len(pool.Buys) == 0 {
		pool.insertSell(sell)
		return
	}
	firstOutput, secondOutput := sell.close(&pool.Buys[0])
	if firstOutput == nil || secondOutput == nil {
		buy := pool.ejectFirstBuy()
		firstOutput, secondOutput := buy.close(sell)
		if firstOutput == nil || secondOutput == nil {
			pool.insertSell(sell)
			pool.insertBuy(buy)
			return
		}
		pool.MainOutputs = append(pool.MainOutputs, *secondOutput)
		pool.MarketOutputs = append(pool.MarketOutputs, *firstOutput)
		pool.OperateBuy(buy)
		return
	}
	pool.MainOutputs = append(pool.MainOutputs, *firstOutput)
	pool.MarketOutputs = append(pool.MarketOutputs, *secondOutput)
	pool.OperateSell(sell)
}

func (pool *tradePool) insertBuy(buy *trade) {
	currentRatio := float64(buy.Offer) / float64(buy.Recieve)
	for addIndex, checkBuy := range pool.Buys {
		checkRatio := float64(checkBuy.Offer) / float64(checkBuy.Recieve)
		if currentRatio > checkRatio {
			pool.Buys = append(
				pool.Buys[:addIndex+1],
				pool.Buys[addIndex:]...,
			)
			pool.Buys[addIndex] = *buy
			return
		}
	}
	pool.Buys = append(pool.Buys, *buy)
}

func (pool *tradePool) ejectFirstSell() *trade {
	sell := pool.Sells[0]
	pool.Sells = pool.Sells[1:]
	return &sell
}

func (pool *tradePool) OperateBuy(buy *trade) {
	if len(pool.Buys) == 0 {
		pool.insertBuy(buy)
		return
	}
	firstOutput, secondOutput := buy.close(&pool.Sells[0])
	if firstOutput == nil || secondOutput == nil {
		sell := pool.ejectFirstSell()
		firstOutput, secondOutput := sell.close(buy)
		if firstOutput == nil || secondOutput == nil {
			pool.insertSell(sell)
			pool.insertBuy(buy)
			return
		}
		pool.MainOutputs = append(pool.MainOutputs, *firstOutput)
		pool.MarketOutputs = append(pool.MarketOutputs, *secondOutput)
		pool.OperateSell(sell)
		return
	}
	pool.MainOutputs = append(pool.MainOutputs, *secondOutput)
	pool.MarketOutputs = append(pool.MarketOutputs, *firstOutput)
	pool.OperateBuy(buy)
}
