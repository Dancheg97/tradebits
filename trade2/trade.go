package trade2

type trade struct {
	Offer   uint64
	Recieve uint64
	Adress  []byte
}

func CreateTrade(
	adress []byte,
	revieve uint64,
	offer uint64,
) *trade {
	if offer == 0 || revieve == 0 || adress == nil {
		return nil
	}
	return &trade{
		Offer:   offer,
		Recieve: revieve,
		Adress:  adress,
	}
}

func (first trade) close(second *trade) bool {
	if first.Offer >= second.Recieve && first.Recieve >= second.Offer {
		firstRatio := float64(first.Offer) / float64(first.Recieve)
		first.Offer = first.Offer - second.Recieve
		first.Recieve = first.Recieve - second.Offer
		secondRatio := float64(first.Offer) / float64(first.Recieve)
		if secondRatio > firstRatio {
			return true
		}
	}
	return false
}
