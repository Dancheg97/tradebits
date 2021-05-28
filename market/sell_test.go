package market

import "testing"

func TestSellMatching(t testing.T) {
	sell := Sell{
		OfferMarket: 100,
		RecieveMain: 50,
	}
	buy := Buy{
		OfferMain:     100,
		RecieveMarket: 50,
	}
	
}
