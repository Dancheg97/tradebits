package trade

// after creation this trade should be attached to some user, then to trade
// pool of some market
type Buy struct {
	Adress  []byte
	Offer   uint64
	Recieve uint64
}

func (b *Buy) close() {
	b.Offer = 0
	b.Recieve = 0
}

func (b *Buy) reduceOffer(amount uint64) {
	if b.Offer != 0 {
		b.Offer = b.Offer - amount
	}
}

func (b *Buy) reduceRecieve(amount uint64) {
	if b.Recieve != 0 {
		b.Recieve = b.Recieve - amount
	}
}

// this function is gonna add single buy offer and operate it for all currently
// matching sell operations
func (t *TradePool) OperateBuy(buy Buy) {
	if len(t.Sells) == 0 {
		t.insertBuy(buy)
		return
	}
	outputs := match(&buy, &t.Sells[0])
	if outputs == nil {
		t.insertBuy(buy)
		return
	}
	t.Outputs = append(t.Outputs, outputs...)
	if t.Sells[0].Offer == 0 {
		t.Sells = t.Sells[1:]
	}
	if buy.Offer == 0 {
		return
	}
	t.OperateBuy(buy)
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
