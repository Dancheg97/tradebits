package trade

import (
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

