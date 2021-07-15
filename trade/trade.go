package trade

type TradePool struct {
	Buys    []Buy
	Sells   []Sell
	Outputs []Output
}

// after creation this trade should be attached to some user, then to trade
// pool of some market
type Buy struct {
	Adress   []byte
	Offer    uint64
	Recieve  uint64
	Attached bool
}

// after creation this trade should be attached to some user, then to trade
// pool of some market
type Sell struct {
	Adress   []byte
	Offer    uint64
	Recieve  uint64
	Attached bool
}

// this struct is used only to transfer data about market outputs for some user
type Output struct {
	Adress []byte
	IsMain bool
	Amount uint64
}

// all trades are alwayts closing to the side better side
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
		b.Offer = 0
		s.Offer = 0
		return []Output{buyOut, sellOut}
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
			buyOutput := Output{
				Adress: b.Adress,
				IsMain: false,
				Amount: b.Recieve,
			}
			sellOutput := Output{
				Adress: s.Adress,
				IsMain: true,
				Amount: b.Offer,
			}
			b.Offer = 0
			s.Offer = potenSellOffer
			s.Recieve = potenSellRecieve
			return []Output{buyOutput, sellOutput}
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
		buyOutput := Output{
			Adress: b.Adress,
			IsMain: false,
			Amount: s.Offer,
		}
		sellOutput := Output{
			Adress: s.Adress,
			IsMain: true,
			Amount: s.Recieve,
		}
		s.Offer = 0
		b.Offer = potentialBuyOffer
		b.Recieve = potentialBuyRecieve
		return []Output{buyOutput, sellOutput}
	}
	return nil
}

// this function is gonna add single buy offer and operate it for all currently
// matching sell operations
func (t *TradePool) OperateBuy(b Buy) {
	if len(t.Sells) == 0 {
		t.insertBuy(b)
		return
	}
	outputs := b.match(&t.Sells[0])
	if outputs == nil {
		t.insertBuy(b)
		return
	}
	t.Outputs = append(t.Outputs, outputs...)
	if t.Sells[0].Offer == 0 {
		t.Sells = t.Sells[1:]
	}
	if b.Offer == 0 {
		return
	}
	t.OperateBuy(b)
}

// this function is made to insert buy to a place, where it should be
// depending on the ratio
func (t *TradePool) insertBuy(b Buy) {
	currentRatio := float64(b.Offer) / float64(b.Recieve)
	for addIndex, checkBuy := range t.Buys {
		checkRatio := float64(checkBuy.Offer) / float64(checkBuy.Recieve)
		if currentRatio > checkRatio {
			t.Buys = append(t.Buys[:addIndex+1], t.Buys[addIndex:]...)
			t.Buys[addIndex] = b
			return
		}
	}
	t.Buys = append(t.Buys, b)
}

// this function is gonna add single sell offer and operate it for all currently
// matching buy operations
func (t *TradePool) OperateSell(s Sell) {
	if len(t.Buys) == 0 {
		t.insertSell(s)
		return
	}
	outputs := t.Buys[0].match(&s)
	if outputs == nil {
		t.insertSell(s)
		return
	}
	t.Outputs = append(t.Outputs, outputs...)
	if t.Buys[0].Offer == 0 {
		t.Buys = t.Buys[1:]
	}
	if s.Offer == 0 {
		return
	}
	t.OperateSell(s)
}

// this function is made to insert sell to a place, where it should be
// depending on the ratio
func (t *TradePool) insertSell(s Sell) {
	currentRatio := float64(s.Offer) / float64(s.Recieve)
	for addIndex, checkSell := range t.Sells {
		checkRatio := float64(checkSell.Offer) / float64(checkSell.Recieve)
		if currentRatio > checkRatio {
			t.Sells = append(t.Sells[:addIndex+1], t.Sells[addIndex:]...)
			t.Sells[addIndex] = s
			return
		}
	}
	t.Sells = append(t.Sells, s)
}
