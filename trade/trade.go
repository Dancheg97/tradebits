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

// all trades are alwayts closing to the side better side
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
	if b.Offer < s.Recieve {
		// that is potential buy closing
		curSellRatio := float64(s.Offer) / float64(s.Recieve)
		potenSellOffer := s.Offer - b.Recieve
		if potenSellOffer > s.Offer {
			return nil
		}
		potenSellRecieve := s.Recieve - b.Offer
		newSellRatio := float64(potenSellOffer) / float64(potenSellRecieve)
		if newSellRatio <= curSellRatio {
			buyOutut := output{
				Adress: b.Adress,
				IsMain: false,
				Amount: b.Recieve,
			}
			sellOutput := output{
				Adress: s.Adress,
				IsMain: true,
				Amount: b.Offer,
			}
			b.Offer = 0
			b.Recieve = 0
			s.Offer = potenSellOffer
			s.Recieve = potenSellRecieve
			return []output{buyOutut, sellOutput}
		}
	}
	// that is potential sell closing
	curBuyRatio := float64(b.Offer) / float64(b.Recieve)
	
	return nil
}
