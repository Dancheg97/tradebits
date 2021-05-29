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

func (new Trade) close(old Trade) (Trade, output, output) {
	firstOutput := output{
		Adress:  new.Adress,
		MainOut: new.Recieve,
	}
	secondOutput := output{
		Adress:    old.Adress,
		MarketOut: new.Offer,
	}
	old.Offer = old.Offer - new.Recieve   // TODO check that
	old.Recieve = old.Recieve - new.Offer // TODO check that
	return old, firstOutput, secondOutput
}
