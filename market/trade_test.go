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
	new := Trade{
		IsSell:  true,
		Offer:   200,
		Recieve: 100,
	}
	old := Trade{
		IsSell:  false,
		Offer:   700,
		Recieve: 400,
	}
	if new.match(old) {
		if new.compare(old) {
			trade, firstOut, secondOut := new.close(old)
			if trade.IsSell != false {
				t.Error("output trade should be buy")
			}
			if firstOut.MainOut != 100 {
				t.Error("first output should be 100")
			}
			if secondOut.MarketOut != 200 {
				t.Error("second output should be 200")
			}
			if trade.Recieve != 200 {
				t.Error("new trade revieve should be 200")
			}
			if trade.Offer != 600 {
				t.Error("new trade offer should be 600")
			}
			return
		}
	}
	t.Error("trade didn't even start")
}

func TestCloseOutput(t *testing.T) {
	new := Trade{
		IsSell:  false,
		Offer:   900,
		Recieve: 300,
	}
	old := Trade{
		IsSell:  true,
		Offer:   250,
		Recieve: 200,
	}
	if new.match(old) {
		if new.compare(old) {
			t.Error("old one should be closing")
		}
		trade, firstOut, secondOut := new.close(old)
		if trade.IsSell == false {
			t.Error("trade should be sell")
		}
		
	}
	t.Error("trade didn't even start")
}
