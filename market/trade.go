package market

type Trade struct {
	Adress  []byte
	IsSell  bool
	Offer   uint64
	Recieve uint64
}

func (in Trade) checkMatch(out Trade) bool {
	inRatio := float64(in.Offer) / float64(in.Recieve)
	outRatio := float64(out.Recieve) / float64(out.Offer)
	return inRatio > outRatio
}

func (in Trade) checkCloseInput(out Trade) bool {
	return in.Offer < out.Recieve
}

func (in Trade) closeInput(out Trade) (Trade, output, output) {
	firstOutput := output{
		Adress:  in.Adress,
		MainOut: in.Recieve,
	}
	secondOutput := output{
		Adress:    out.Adress,
		MarketOut: in.Offer,
	}
	out.Offer = out.Offer - in.Recieve   // TODO check that
	out.Recieve = out.Recieve - in.Offer // TODO check that
	return out, firstOutput, secondOutput
}

func (in Trade) closeOutput(out Trade) (Trade, output, output) {
	
}
