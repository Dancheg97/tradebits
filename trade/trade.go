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

// this struct is used only to transfer data about market outputs for some user
type Output struct {
	Adress []byte
	Main   uint64
	Market uint64
}

// all trades are alwayts closing to the side better side
func (b *Buy) match(s *Sell) []Output {
	if float64(b.Offer)/float64(s.Recieve) >= float64(b.Recieve/s.Offer) {
		if s.Offer >= b.Recieve && b.Offer >= s.Recieve {
			buyerOutput := Output{
				Adress: b.Adress,
				Market: b.Recieve,
			}
			sellerOutput := Output{
				Adress: s.Adress,
				Main:   b.Offer,
			}
			s.Offer = s.Offer - b.Recieve
			s.Recieve = s.Recieve - b.Offer
			if s.Recieve == 0 && s.Offer != 0 {
				sellerOutput.Market = s.Offer
				s.Offer = 0
			}
			b.Offer = 0
			b.Recieve = 0
			return []Output{
				buyerOutput,
				sellerOutput,
			}
		}
		buyerOutput := Output{
			Adress: b.Adress,
			Market: s.Offer,
		}
		sellerOutput := Output{
			Adress: s.Adress,
			Main:   s.Recieve,
		}
		b.Offer = b.Offer - s.Recieve
		b.Recieve = b.Recieve - s.Offer
		s.Offer = 0
		s.Recieve = 0
		return []Output{
			buyerOutput,
			sellerOutput,
		}
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
