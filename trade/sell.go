package trade

// after creation this trade should be attached to some user, then to trade
// pool of some market
type Sell struct {
	Adress  []byte
	Offer   uint64
	Recieve uint64
}

func (s *Sell) close() {
	s.Offer = 0
	s.Recieve = 0
}

func (s *Sell) reduceOffer(amount uint64) {
	if s.Offer != 0 {
		s.Offer = s.Offer - amount
	}
}

func (s *Sell) reduceRecieve(amount uint64) {
	if s.Recieve != 0 {
		s.Recieve = s.Recieve - amount
	}
}

// this function is gonna add single sell offer and operate it for all currently
// matching buy operations
func (t *TradePool) OperateSell(s Sell) {
	if len(t.Buys) == 0 {
		t.insertSell(s)
		return
	}
	outputs := match(&t.Buys[0], &s)
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
