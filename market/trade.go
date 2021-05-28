package market

type Trade struct {
	Adress  []byte
	IsSell  bool
	Offer   uint64
	Recieve uint64
}

type output struct {
	Adress    []byte
	MainOut   uint64
	MarketOut uint64
}

func (x Trade) checkMatch(y Trade) bool {
	inRatio := float64(x.Offer)/float64(x.Recieve)
	outRatio := float64(y.Recieve)/float64(y.Offer)
	return inRatio > outRatio
}

// func (x Trade) inpClose(y Trade) bool {
// 	return 
// }