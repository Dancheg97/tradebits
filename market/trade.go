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
		if new.IsSell {
			newOutput := output{
				Adress:    new.Adress,
				MainOut:   new.Offer - potentialNewOffer,
				MarketOut: new.Recieve,
			}
			oldOutput := output{
				Adress:  old.Adress,
				MainOut: potentialNewOffer,
			}
			old.Offer = old.Offer - new.Recieve
			old.Recieve = old.Recieve - potentialNewOffer
			return true, []Trade{old}, []output{newOutput, oldOutput}
		}
		newOutput := output{
			Adress:    new.Adress,
			MarketOut: new.Offer - potentialNewOffer,
			MainOut:   new.Recieve,
		}
		oldOutput := output{
			Adress:    old.Adress,
			MarketOut: potentialNewOffer,
		}
		old.Offer = old.Offer - new.Recieve
		old.Recieve = old.Recieve - potentialNewOffer
		return true, []Trade{old}, []output{newOutput, oldOutput}
	}
	newRatio := float64(new.Recieve) / float64(new.Offer)
	oldRatio := float64(old.Offer) / float64(old.Recieve)
	if newRatio < oldRatio {
		return false, []Trade{new, old}, []output{}
	}
	if new.IsSell {
		newOutput := output{
			Adress:  new.Adress,
			MainOut: old.Offer,
		}
		oldOutput := output{
			Adress:    old.Adress,
			MarketOut: old.Recieve,
		}
		new.Offer = new.Offer - old.Recieve
		new.Recieve = new.Recieve - old.Offer
		return true, []Trade{new}, []output{newOutput, oldOutput}
	}
	newOutput := output{
		Adress:    new.Adress,
		MarketOut: old.Offer,
	}
	oldOutput := output{
		Adress:  old.Adress,
		MainOut: old.Recieve,
	}
	new.Offer = new.Offer - old.Recieve
	new.Recieve = new.Recieve - old.Offer
	return true, []Trade{new}, []output{newOutput, oldOutput}
}
