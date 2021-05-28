package asset

import "math"

type Buy struct {
	Adress       []byte
	OfferMain    uint64
	RecieveAsset uint64
}

type Sell struct {
	Adress      []byte
	OfferAsset  uint64
	RecieveMain uint64
}

type Output struct {
	Adress   []byte
	MainOut  uint64
	AssetOut uint64
}

func S2BCheckMatch(sell Sell, buy Buy) bool {
	sellRatio := float64(buy.RecieveAsset / buy.OfferMain)
	buyRatio := float64(sell.OfferAsset / sell.RecieveMain)
	return sellRatio < buyRatio
}

func S2BIfCloseSeller(sell Sell, buy Buy) bool {
	return sell.RecieveMain < buy.OfferMain
}

func S2BCloseSeller(sell Sell, buy Buy) (Buy, Output, Output) {
	ratio := float64(buy.RecieveAsset) / float64(buy.OfferMain)
	upRoundOut := uint64(math.Ceil(ratio * float64(sell.RecieveMain)))
	sellerOutput := Output{
		AssetOut: sell.OfferAsset - upRoundOut,
		MainOut:  sell.RecieveMain,
	}
	buyerOutput := Output{
		AssetOut: upRoundOut,
	}
	buy.OfferMain = buy.OfferMain - sell.RecieveMain
	buy.RecieveAsset = buy.RecieveAsset - upRoundOut
	return buy, sellerOutput, buyerOutput
}

func S2BCloseBuyer(sell Sell, buy Buy) (Sell, Output, Output) {
	sellerOutput := Output{
		MainOut: buy.OfferMain,
	}
	buyerOutput := Output{
		AssetOut: buy.RecieveAsset,
	}
	sell.OfferAsset = sell.OfferAsset - buy.RecieveAsset
	sell.RecieveMain = sell.RecieveMain - buy.OfferMain
	return sell, sellerOutput, buyerOutput
}

func B2SCheckMatch(buy Buy, sell Sell) bool {
	sellRatio := float64(buy.RecieveAsset / buy.OfferMain)
	buyRatio := float64(sell.OfferAsset / sell.RecieveMain)
	return sellRatio < buyRatio
}
