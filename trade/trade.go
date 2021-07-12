package trade

type TradePool struct {
	Buys    []Buy
	Sells   []Sell
	Outputs []Output
}

// after creation this trade should be attached to some user, then to trade
// pool of some market
type Buy struct {
	Adress  []byte
	Offer   uint64
	Recieve uint64
}

// after creation this trade should be attached to some user, then to trade
// pool of some market
type Sell struct {
	Adress  []byte
	Offer   uint64
	Recieve uint64
}

type Output struct {
	Adress []byte
	IsMain bool
	Amount uint64
}

func (b *Buy) match(s *Sell) []Output {
	if b.Offer == s.Recieve && b.Recieve == s.Offer {
		buyOut := Output{
			Adress: b.Adress,
			IsMain: false,
			Amount: s.Offer,
		}
		sellOut := Output{
			Adress: s.Adress,
			IsMain: true,
			Amount: b.Offer,
		}
		return []Output{buyOut, sellOut}
	}
	return nil
}
