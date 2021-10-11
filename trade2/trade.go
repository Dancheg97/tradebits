package trade2

type trade struct {
	Offer   uint64
	Recieve uint64
	Adress  []byte
}

type output struct {
	Adress []byte
	Amount uint64
}

func CreateTrade(
	adress []byte,
	offer uint64,
	recieve uint64,
) *trade {
	if offer == 0 || recieve == 0 || adress == nil {
		return nil
	}
	return &trade{
		Offer:   offer,
		Recieve: recieve,
		Adress:  adress,
	}
}

// first trade is closing second trade
func (first *trade) close(second *trade) (*output, *output) {
	if first.Offer >= second.Recieve && first.Recieve >= second.Offer {
		firstRatio := float64(first.Offer) / float64(first.Recieve)
		newFirstOffer := first.Offer - second.Recieve
		newfirstRecieve := first.Recieve - second.Offer
		secondRatio := float64(newFirstOffer) / float64(newfirstRecieve)
		if secondRatio > firstRatio {
			firstOutput := output{
				Adress: first.Adress,
				Amount: second.Offer,
			}
			secondOutput := output{
				Adress: second.Adress,
				Amount: second.Recieve,
			}
			return &firstOutput, &secondOutput
		}
	}
	return nil, nil
}
