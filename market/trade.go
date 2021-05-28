package market

import "math"

type Buy struct {
	Adress        []byte
	OfferMain     uint64
	RecieveMarket uint64
}

type Sell struct {
	Adress      []byte
	OfferMarket uint64
	RecieveMain uint64
}

type Output struct {
	Adress    []byte
	MainOut   uint64
	MarketOut uint64
}

func S2BCheckMatch(sell Sell, buy Buy) bool {
	sellRatio := float64(buy.RecieveMarket / buy.OfferMain)
	buyRatio := float64(sell.OfferMarket / sell.RecieveMain)
	return sellRatio < buyRatio
}

func S2BIfCloseSeller(sell Sell, buy Buy) bool {
	return sell.RecieveMain < buy.OfferMain
}

func S2BCloseSeller(sell Sell, buy Buy) (Buy, Output, Output) {
	ratio := float64(buy.RecieveMarket) / float64(buy.OfferMain)
	upRoundOut := uint64(math.Ceil(ratio * float64(sell.RecieveMain)))
	sellerOutput := Output{
		MarketOut: sell.OfferMarket - upRoundOut,
		MainOut:   sell.RecieveMain,
	}
	buyerOutput := Output{
		MarketOut: upRoundOut,
	}
	buy.OfferMain = buy.OfferMain - sell.RecieveMain
	buy.RecieveMarket = buy.RecieveMarket - upRoundOut
	return buy, sellerOutput, buyerOutput
}

func S2BCloseBuyer(sell Sell, buy Buy) (Sell, Output, Output) {
	sellerOutput := Output{
		MainOut: buy.OfferMain,
	}
	buyerOutput := Output{
		MarketOut: buy.RecieveMarket,
	}
	sell.OfferMarket = sell.OfferMarket - buy.RecieveMarket
	sell.RecieveMain = sell.RecieveMain - buy.OfferMain
	return sell, sellerOutput, buyerOutput
}

func B2SCheckMatch(buy Buy, sell Sell) bool {
	sellRatio := float64(buy.RecieveMarket / buy.OfferMain)
	buyRatio := float64(sell.OfferMarket / sell.RecieveMain)
	return sellRatio < buyRatio
}

func B2SIfCloseBuyer(buy Buy, sell Sell) bool {
	return buy.RecieveMarket < sell.OfferMarket
}
