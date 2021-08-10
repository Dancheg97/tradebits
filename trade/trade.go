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
func (buy *Buy) match(sell *Sell) []Output {
	if float64(buy.Offer)/float64(sell.Recieve) >= float64(buy.Recieve/sell.Offer) {
		buyerOutput := Output{
			Adress: buy.Adress,
		}
		sellerOutput := Output{
			Adress: sell.Adress,
		}
		if buy.Offer < sell.Recieve {
			defer buy.close()
			defer sell.reduceOffer(buy.Recieve)
			defer sell.reduceRecieve(buy.Offer)
			buyerOutput.Market = buy.Recieve
			sellerOutput.Main = buy.Offer
		}
		if sell.Offer < buy.Recieve {
			defer sell.close()
			defer buy.reduceOffer(sell.Recieve)
			defer buy.reduceRecieve(sell.Offer)
			buyerOutput.Market = sell.Offer
			sellerOutput.Main = sell.Recieve
		}
		return []Output{
			buyerOutput,
			sellerOutput,
		}
	}
	return nil
}

func (b *Buy) close() {
	b.Offer = 0
	b.Recieve = 0
}

func (s *Sell) close() {
	s.Offer = 0
	s.Recieve = 0
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
