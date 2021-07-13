package trade

import (
	"testing"
)

func TestMatchSameValues(t *testing.T) {
	buy := Buy{
		Adress:  []byte{0},
		Offer:   100,
		Recieve: 50,
	}
	sell := Sell{
		Adress:  []byte{1},
		Offer:   50,
		Recieve: 100,
	}
	outputs := buy.match(&sell)
	if len(outputs) != 2 {
		t.Error("there should be 2 outputs for both users")
	}
	buyerOutput := outputs[0]
	if buyerOutput.Amount != 50 {
		t.Error("buyers output should be equal to 50")
	}
	if buyerOutput.IsMain {
		t.Error("first output should be market")
	}
	sellerOutput := outputs[1]
	if sellerOutput.Amount != 100 {
		t.Error("seller output should be equal to zero")
	}
	if !sellerOutput.IsMain {
		t.Error("seller should recieve main coin")
	}
}

func TestMatchCancelByRatio(t *testing.T) {
	buy := Buy{
		Adress:  []byte{0},
		Offer:   10,
		Recieve: 20,
	}
	sell := Sell{
		Adress:  []byte{1},
		Offer:   10,
		Recieve: 20,
	}
	output := buy.match(&sell)
	if len(output) != 0 {
		t.Error("this trades should not match")
	}
}

func TestMatchBuyClosing(t *testing.T) {
	
}

func TestMatchSellClosing(t *testing.T) {

}
