package asset

import "math"

type ProcessTrade interface {
	trade(sell Sell, buy Buy)
}

// request to buy some asset
type Buy struct {
	Adress       []byte
	OfferMain    uint64
	RecieveAsset uint64
}

// request to sell some asset
type Sell struct {
	Adress      []byte
	OfferAsset  uint64
	RecieveMain uint64
}

// struct containing info about outputs
type Output struct {
	Adress   []byte
	MainOut  uint64
	AssetOut uint64
}

func CheckMatch(sell Sell, buy Buy) bool {
	return float64(buy.RecieveAsset/buy.OfferMain) < float64(sell.OfferAsset/sell.RecieveMain)
}

func IfCloseSeller(sell Sell, buy Buy) bool {
	return sell.RecieveMain < buy.OfferMain
}

// first output is for seller, second output is for buyer
func CloseSeller(sell Sell, buy Buy) (Buy, Output, Output) {
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

// first output is for seller, second output is for buyer
func CloseBuyer(sell Sell, buy Buy) (Sell, Output, Output) {
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

func tradeSellToBuy(sell Sell, buy Buy) ()