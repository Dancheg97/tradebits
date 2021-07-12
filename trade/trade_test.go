package trade

import (
	"reflect"
	"testing"
)

func TestBuyCreation(t *testing.T) {
	buy1 := NewBuy([]byte{0}, 1000, 1000)
	if buy1.Offer != 1000 {
		t.Error("buy offer should be equal to 1000")
	}
	buy2 := NewBuy([]byte{0}, 0, 1000)
	if buy2 != nil {
		t.Error("buy 2 should not be created due to 0 offer")
	}
}

func TestSellCreation(t *testing.T) {
	sell1 := NewSell([]byte{0}, 1000, 1000)
	if sell1.Offer != 1000 {
		t.Error("sell offer should be equal to 1000")
	}
	sell2 := NewSell([]byte{1}, 0, 1000)
	if sell2 != nil {
		t.Error("sell2 should not be created with 0 offer")
	}
}

func TestMatchSameValues(t *testing.T) {
	buy := NewBuy([]byte{0}, 10, 20)
	sell := NewSell([]byte{1}, 20, 10)
	rez := buy.match(sell)
	for _, output := range rez {
		if reflect.DeepEqual(output.Adress, []byte{0}) {
			if output.IsMain {
				t.Error("buyer output should be for market")
			}
			if output.Amount != 10 {
				
			}
		}
	}
}
