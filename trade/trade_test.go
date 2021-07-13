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
	buy := Buy{
		Adress:  []byte{0},
		Offer:   40,
		Recieve: 90,
	}
	sell := Sell{
		Adress:  []byte{1},
		Offer:   100,
		Recieve: 50,
	}
	output := buy.match(&sell)
	if len(output) != 2 {
		t.Error("there should be 2 outputs in this trade")
	}
	buyOutput := output[0]
	if buyOutput.IsMain {
		t.Error("buyers output should be for mkt asset")
	}
	if buyOutput.Amount != 90 {
		t.Error("buyers output should be equal to 90")
	}
	sellOutput := output[1]
	if !sellOutput.IsMain {
		t.Error("seller output should be for main coin")
	}
	if sellOutput.Amount != 40 {
		t.Error("seller output should be equal to 40")
	}
	if buy.Offer != 0 {
		t.Error("new buyers offer should be equal to 0, this trade should be closed")
	}
	if sell.Offer != 10 {
		t.Error("new sell offer should be equal to 10")
	}
	if sell.Recieve != 10 {
		t.Error("new sell recieve should be equal to 10")
	}
}

func TestMatchSellClosing(t *testing.T) {
	buy := Buy{
		Adress:  []byte{0},
		Offer:   100,
		Recieve: 50,
	}
	sell := Sell{
		Adress:  []byte{1},
		Offer:   45,
		Recieve: 49,
	}
	outputs := buy.match(&sell)
	if len(outputs) != 2 {
		t.Error("Length of outputs should be equal to 2")
	}
	buyOutput := outputs[0]
	if buyOutput.IsMain {
		t.Error("buyers output should be for market asset")
	}
	if buyOutput.Amount != 45 {
		t.Error("buyer output should be equal to 45")
	}
	sellOutput := outputs[1]
	if !sellOutput.IsMain {
		t.Error("seller output should be for main asset")
	}
	if sellOutput.Amount != 49 {
		t.Error("seller should get 49 output")
	}
	if sell.Offer != 0 {
		t.Error("seller output should be equal to zero")
	}
	if buy.Offer != 51 {
		t.Error("new buyer output should be equal to 51")
	}
	if buy.Recieve != 5 {
		t.Error("new buyer recieve should be equal to 5")
	}
}
