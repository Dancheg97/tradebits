package market

import "math"

type Trade struct {
	Adress  []byte
	IsSell  bool
	Offer   uint64
	Recieve uint64
}

func (new Trade) operate(old Trade) ([]Trade, []output) {
	if new.Recieve < old.Offer {
		ratio := float64(old.Recieve) / float64(old.Offer)
		potentialNewOffer := uint64(math.Ceil(float64(new.Recieve) * ratio))
		if potentialNewOffer > new.Offer {
			return []Trade{new, old}, nil
		}
		newOutput := output{Adress: new.Adress}
		oldOutput := output{Adress: old.Adress}
		if old.IsSell {
			newOutput.MainOut = new.Offer - potentialNewOffer
			newOutput.MarketOut = new.Recieve
			oldOutput.MainOut = potentialNewOffer
		} else {
			newOutput.MarketOut = new.Offer - potentialNewOffer
			newOutput.MainOut = new.Recieve
			oldOutput.MarketOut = potentialNewOffer
		}
		old.Offer = old.Offer - new.Recieve // add dot after old in var name
		old.Recieve = old.Recieve - potentialNewOffer
		return []Trade{old}, []output{newOutput, oldOutput}
	}
	newRatio := float64(new.Recieve) / float64(new.Offer)
	oldRatio := float64(old.Offer) / float64(old.Recieve)
	if newRatio > oldRatio {
		return []Trade{new, old}, nil
	}
	newOutput := output{Adress: new.Adress}
	oldOutput := output{Adress: old.Adress}
	if new.IsSell {
		newOutput.MainOut = old.Offer
		oldOutput.MarketOut = old.Recieve
	} else {
		newOutput.MarketOut = old.Offer
		oldOutput.MainOut = old.Recieve
	}
	new.Offer = new.Offer - old.Recieve
	new.Recieve = new.Recieve - old.Offer
	if new.Offer == 0 && new.Recieve == 0 {
		return []Trade{}, []output{newOutput, oldOutput}
	}
	return []Trade{new}, []output{newOutput, oldOutput}
}

// assistive func to add trade to curr trade list in proper place by ratio
func (m *market) addTrade(t Trade) {
	currRatio := float64(t.Offer) / float64(t.Recieve)
	if t.IsSell {
		if len(m.Sells) == 0 {
			m.Sells = append(m.Sells, t)
			return
		}
		for index, sell := range m.Sells {
			sellRatio := float64(sell.Offer) / float64(sell.Recieve)
			if currRatio > sellRatio {
				m.Sells = append(m.Sells[:index+1], m.Sells[index:]...)
				m.Sells[index] = t
			}
		}
	} else {
		if len(m.Buys) == 0 {
			m.Buys = append(m.Buys, t)
			return
		}
		for index, buy := range m.Buys {
			buyRatio := float64(buy.Offer) / float64(buy.Recieve)
			if currRatio > buyRatio {
				m.Buys = append(m.Buys[:index+1], m.Buys[index:]...)
				m.Buys[index] = t
			}
		}
	}
}
