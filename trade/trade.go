package trade

type TradePool struct {
	Buys    []Buy
	Sells   []Sell
	Outputs []output
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

// this struct is used only to transfer data about market outputs for some user
type output struct {
	Adress []byte
	IsMain bool
	Amount uint64
}

func (b *Buy) match(s *Sell) []output {
	if b.Offer == s.Recieve && b.Recieve == s.Offer {
		buyOut := output{
			Adress: b.Adress,
			IsMain: false,
			Amount: s.Offer,
		}
		sellOut := output{
			Adress: s.Adress,
			IsMain: true,
			Amount: b.Offer,
		}
		return []output{buyOut, sellOut}
	}
	return nil
}
