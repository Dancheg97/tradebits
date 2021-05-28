package market

import "testing"

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
	firstMatch := t1.checkMatch(t2)
	secondMatch := t2.checkMatch(t1)
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
	firstMatch := t1.checkMatch(t2)
	secondMatch := t2.checkMatch(t1)
	firstClose := t1.checkCloseInput(t2)
	secondClose := t2.checkCloseInput(t1)
	if firstMatch && secondMatch {
		if firstClose || secondClose {
			return
		}
	}
	t.Error("trades do not match")
}

