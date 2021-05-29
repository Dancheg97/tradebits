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
	firstClose := t1.isLessThan(t2)
	secondClose := t2.isLessThan(t1)
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
		Offer:   201,
		Recieve: 300,
	}
	buy := Trade{
		Adress:  []byte("new seller"),
		IsSell:  false,
		Offer:   700,
		Recieve: 400,
	}
	if sell.match(buy) {
		if sell.isLessThan(buy) {
			trade, firstOut, secondOut := sell.closingBy(buy)
			if !reflect.DeepEqual(trade.Adress, []byte("new seller")) {
				t.Error("left trade should be on seller")
			}
			if trade.Offer != 600 {
				t.Error("trade offer should be 600")
			}
			if trade.Recieve != 200 {
				t.Error("trade recieve should be 200")
			}
			if trade.IsSell != false {
				t.Error("trade should be sell")
			}
			if reflect.DeepEqual(firstOut.Adress, []byte("new seller")) {
				t.Error("second output should be for seller")
			}
			if firstOut.MainOut != 100 {
				t.Error("Output for seller should be 100 market")
			}
			if reflect.DeepEqual(secondOut.Adress, []byte("old buyer")) {
				t.Error("first output should be for buyer")
			}
			if secondOut.MarketOut != 200 {
				t.Error("output for buyer should be 200 market")
			}
			return
		}
	}
	t.Error("trade didn't even start")
}

func TestAttemptToGetBigUINT(t *testing.T) {
	sell := Trade{
		Adress:  []byte("seller"),
		IsSell:  true,
		Offer:   500,
		Recieve: 300,
	}
	buy := Trade{
		Adress:  []byte("buyer"),
		IsSell:  false,
		Offer:   700,
		Recieve: 300,
	}
	isProcessed, trade, output1, output2 := sell.processTrade(buy)
	t.Error(isProcessed)
	t.Error(trade)
	t.Error(output1)
	t.Error(output2)
}
