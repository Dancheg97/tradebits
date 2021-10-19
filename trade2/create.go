package trade2

type tradePool struct {
	Buys          []*trade
	MainOutputs   []*output
	Sells         []*trade
	MarketOutputs []*output
}

func CreatePool() *tradePool {
	return &tradePool{
		Buys:  []*trade{},
		Sells: []*trade{},
	}
}