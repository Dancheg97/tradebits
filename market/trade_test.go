package market

import (
	"reflect"
	"testing"
)

func TestCheckMatching(t *testing.T) {
	t1 := Trade{
		IsSell:  true,
		Offer:   200,
		Recieve: 100,
	}
	t2 := Trade{
		IsSell:  false,
		Offer:   200,
		Recieve: 100,
	}
	firstMatch := t1.match(t2)
	secondMatch := t2.match(t1)
	if firstMatch && secondMatch {
		return
	}
	t.Error("trades do not match")
}

func TestCloseCheck(t *testing.T) {
	t1 := Trade{
		IsSell:  true,
		Offer:   200,
		Recieve: 100,
	}
	t2 := Trade{
		IsSell:  false,
		Offer:   800,
		Recieve: 400,
	}
	firstMatch := t1.match(t2)
	secondMatch := t2.match(t1)
	firstClose := t1.compare(t2)
	secondClose := t2.compare(t1)
	if firstMatch && secondMatch {
		if firstClose || secondClose {
			return
		}
	}
	t.Error("trades do not match")
}

func TestCloseInput(t *testing.T) {
	sell := Trade{
		Adress:  []byte("old buyer"),
		IsSell:  true,
		Offer:   200,
		Recieve: 100,
	}
	buy := Trade{
		Adress:  []byte("new seller"),
		IsSell:  false,
		Offer:   700,
		Recieve: 400,
	}
	if sell.match(buy) {
		if sell.compare(buy) {
			trade, firstOut, secondOut := sell.close(buy)
			if !reflect.DeepEqual(trade.Adress, []byte("new seller")) {
				t.Error("left trade should be on seller")
			}
			if trade.Offer != 200 {
				t.Error("trade offer should be 200")
			}
			if trade.Recieve != 600 {
				t.Error("trade recieve should be 100")
			}
			if trade.IsSell != false {
				t.Error("trade should be sell")
			}
			if reflect.DeepEqual(firstOut.Adress, []byte("old buyer")) {
				t.Error("first output should be for buyer")
			}
			if firstOut.MainOut != 100 {
				t.Error("output for buyer should be 100")
			}
			if reflect.DeepEqual(secondOut.Adress, []byte("new seller")) {
				t.Error("second output should be for seller")
			}
			if secondOut.MarketOut != 200 {
				t.Error("Output ")
			}
			return
		}
	}
	t.Error("trade didn't even start")
}

func TestCloseOutput(t *testing.T) {

}
