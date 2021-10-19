package trade2

func (pool *tradePool) OperateSell(sell *trade) {
	if len(pool.Buys) == 0 {
		pool.insertSell(sell)
		return
	}
	firstOutput, secondOutput := sell.close(pool.Buys[0])
	if firstOutput == nil || secondOutput == nil {
		buy := pool.ejectFirstBuy()
		firstOutput, secondOutput := buy.close(sell)
		if firstOutput == nil || secondOutput == nil {
			pool.insertSell(sell)
			pool.insertBuy(buy)
			return
		}
		pool.MainOutputs = append(pool.MainOutputs, secondOutput)
		pool.MarketOutputs = append(pool.MarketOutputs, firstOutput)
		pool.OperateBuy(buy)
		return
	}
	pool.MainOutputs = append(pool.MainOutputs, firstOutput)
	pool.MarketOutputs = append(pool.MarketOutputs, secondOutput)
	pool.OperateSell(sell)
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
			pool.Sells[addIndex] = sell
			return
		}
	}
	pool.Sells = append(pool.Sells, sell)
}

func (pool *tradePool) ejectFirstSell() *trade {
	sell := pool.Sells[0]
	pool.Sells = pool.Sells[1:]
	return sell
}
