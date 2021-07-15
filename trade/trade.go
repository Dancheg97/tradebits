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
		curSellRatio := float64(s.Offer) / float64(s.Recieve)
		potenSellOffer := s.Offer - b.Recieve
		if potenSellOffer > s.Offer {
			return nil
		}
		potenSellRecieve := s.Recieve - b.Offer
		newSellRatio := float64(potenSellOffer) / float64(potenSellRecieve)
		if newSellRatio <= curSellRatio {
			buyOutput := output{
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
			return []output{buyOutput, sellOutput}
		}
		return nil
	}
	curBuyRatio := float64(b.Offer) / float64(b.Recieve)
	potentialBuyOffer := b.Offer - s.Recieve
	potentialBuyRecieve := b.Recieve - s.Offer
	if potentialBuyRecieve > b.Recieve {
		return nil
	}
	newBuyRatio := float64(potentialBuyOffer) / float64(potentialBuyRecieve)
	if newBuyRatio >= curBuyRatio {
		buyOutput := output{
			Adress: b.Adress,
			IsMain: false,
			Amount: s.Offer,
		}
		sellOutput := output{
			Adress: s.Adress,
			IsMain: true,
			Amount: s.Recieve,
		}
		s.Offer = 0
		s.Recieve = 0
		b.Offer = potentialBuyOffer
		b.Recieve = potentialBuyRecieve
		return []output{buyOutput, sellOutput}
	}
	return nil
}

func (t *TradePool) AddBuy(b Buy) {
	if len(t.Sells) == 0 {
		t.Buys = []Buy{b}
		return
	}
	outputs := b.match(&t.Sells[0])
	if outputs == nil {
		currentRatio := float64(b.Offer) / float64(b.Recieve)
		for addIndex, checkBuy := range t.Buys {
			checkRatio := float64(checkBuy.Offer) / float64(checkBuy.Recieve)
			if currentRatio > checkRatio {
				t.Buys = append(t.Buys[:addIndex], t.Buys[addIndex-1:]...)
				t.Buys[addIndex] = b
				return
			}
			t.Buys = append(t.Buys, b)
		}
		return
	}
	t.Outputs = append(t.Outputs, outputs...)
	if t.Sells[0].Offer == 0 {
		t.Sells = t.Sells[1:]
	}
	if b.Offer == 0 {
		return
	}
	t.AddBuy(b)
}

func (t *TradePool) AddSell(s Sell) {
	if len(t.Buys) == 0 {
		t.Sells = []Sell{s}
		return
	}
}
