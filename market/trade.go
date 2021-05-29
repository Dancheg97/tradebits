package market

type Trade struct {
	Adress  []byte
	IsSell  bool
	Offer   uint64
	Recieve uint64
}

func (new Trade) match(old Trade) bool {
	newRatio := float64(new.Offer) / float64(new.Recieve)
	oldRatio := float64(old.Recieve) / float64(old.Offer)
	return newRatio > oldRatio
}

func (new Trade) compare(old Trade) bool {
	return new.Offer < old.Recieve
}

/*
Conditions to close trade:
1) Trades should match by ratio
2) One that is left open should increase it's ratio
3) The sum of input for each 'main' and 'market' should be the same as output
*/
func (small Trade) close(big Trade) (Trade, output, output) {
	if small.IsSell {
		newOutput := output{
			Adress:  small.Adress,
			MainOut: small.Recieve,
		}
		oldOutput := output{
			Adress:    big.Adress,
			MarketOut: small.Offer,
		}
		big.Offer = big.Offer - small.Recieve   // TODO check that
		big.Recieve = big.Recieve - small.Offer // TODO check that
		return big, newOutput, oldOutput
	}
	oldOutput := output{
		Adress:  small.Adress,
		MainOut: small.Recieve,
	}
	newOutput := output{
		Adress:    big.Adress,
		MarketOut: small.Offer,
	}
}
