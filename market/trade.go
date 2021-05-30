package market

import "math"

type Trade struct {
	Adress  []byte
	IsSell  bool
	Offer   uint64
	Recieve uint64
}

func (new Trade) operate(old Trade) (bool, []Trade, []output) {
	if new.Recieve < old.Offer {
		ratio := float64(old.Recieve) / float64(old.Offer)
		potentialNewOffer := uint64(math.Ceil(float64(new.Recieve) * ratio))
		if potentialNewOffer > new.Offer {
			return false, []Trade{new, old}, nil
		}
		newOutput := output{Adress: new.Adress}
		oldOutput := output{Adress: old.Adress}
		if new.IsSell {
			newOutput.MainOut = new.Offer - potentialNewOffer
			newOutput.MarketOut = new.Recieve
			oldOutput.MainOut = potentialNewOffer
		} else {
			newOutput.MarketOut = new.Offer - potentialNewOffer
			newOutput.MainOut = new.Recieve
			oldOutput.MarketOut = potentialNewOffer
		}
		old.Offer = old.Offer - new.Recieve
		old.Recieve = old.Recieve - potentialNewOffer
		return true, []Trade{old}, []output{newOutput, oldOutput}
	}
	newRatio := float64(new.Recieve) / float64(new.Offer)
	oldRatio := float64(old.Offer) / float64(old.Recieve)
	if newRatio > oldRatio {
		return false, []Trade{new, old}, []output{}
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
	return true, []Trade{new}, []output{newOutput, oldOutput}
}
