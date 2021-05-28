package market

import "testing"

func TestCheckMatching(t *testing.T) {
	sell := Trade{
		IsSell:  true,
		Offer:   101,
		Recieve: 100,
	}
	buy := Trade{
		IsSell:  false,
		Offer:   101,
		Recieve: 100,
	}
	firstMatch := sell.checkMatch(buy)
	secondMatch := buy.checkMatch(sell)
	if firstMatch && secondMatch {
		return
	}
	t.Error("trades do not match")
}
