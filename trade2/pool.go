package trade2

type pool struct {
	Buys  []trade
	Sells []trade
}

func CreatePool() *pool {
	return &pool{
		Buys:  []trade{},
		Sells: []trade{},
	}
}

