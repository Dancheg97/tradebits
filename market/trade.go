package market

type Trade struct {
	Adress  []byte
	IsSell  bool
	Offer   uint64
	Recieve uint64
}

// checks if 2 trades match by the ratio
func (new Trade) match(old Trade) bool {
	newRatio := float64(new.Offer) / float64(new.Recieve)
	oldRatio := float64(old.Recieve) / float64(old.Offer)
	return newRatio > oldRatio
}

func (new Trade) isLessThan(old Trade) bool {
	return new.Offer < old.Recieve
}

/*
This function is closing the trade on which operation is performed. So the
trade on which operation is performed should be smaller.

Conditions to closingBy trade:
 1) Trades should match by ratio
 2) Output trade should incraese its ratio
 3) The sum of input for each "main" and "market" should be the same as output
*/
func (small Trade) closingBy(big Trade) (Trade, output, output) {
	if small.IsSell {
		firstOutput := output{
			Adress:  small.Adress,
			MainOut: small.Recieve,
		}
		secondOutput := output{
			Adress:    big.Adress,
			MarketOut: small.Offer,
		}
		big.Offer = big.Offer - small.Recieve
		big.Recieve = big.Recieve - small.Offer
		return big, firstOutput, secondOutput
	}
	firstOutput := output{
		Adress:    small.Adress,
		MarketOut: small.Recieve,
	}
	secondOutput := output{
		Adress:  big.Adress,
		MainOut: small.Offer,
	}
	big.Offer = big.Offer - small.Recieve
	big.Recieve = big.Recieve - small.Offer
	return big, firstOutput, secondOutput
}

func (new Trade) processTrade(old Trade) (bool, Trade, output, output) {
	if !new.match(old) {
		return false, new, output{}, output{}
	}
	if new.isLessThan(old) {
		updatedOld, output1, output2 := new.closingBy(old)
		return true, updatedOld, output1, output2
	}
	updatedNew, output1, output2 := old.closingBy(new)
	return true, updatedNew, output1, output2
}
