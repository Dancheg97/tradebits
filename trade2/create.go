package trade

type trade struct {
	Adress  []byte
	Offer   uint64
	Recieve uint64
}

type pool struct {
	Buys  []trade
	Sells []trade
}

func CreateTrade(adress []byte, offer uint64, recieve uint64) trade {
	return trade{
		Adress:  adress,
		Offer:   offer,
		Recieve: recieve,
	}
}

func CreatePool() pool {
	return pool{
		Buys:  []trade{},
		Sells: []trade{},
	}
}
