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

func TestBigRecieveOverSmallOffer(t *testing.T) {
	small := Trade{
		Adress:  []byte("smallSeller"),
		IsSell:  true,
		Offer:   201,
		Recieve: 300,
	}
	big := Trade{
		Adress:  []byte("bigBuyer"),
		IsSell:  false,
		Offer:   700,
		Recieve: 200,
	}
	if small.match(big) {
		if small.compare(big) {
			
		}
	}
	t.Error("trade didn't start")
}
