package trade

import "testing"

func TestAddSingleBuy(t *testing.T) {
	pool := CreatePool()
	buy := CreateTrade([]byte{0}, 1, 1)
	pool.AddBuy(&buy)
	if len(pool.Buys) != 1 {
		t.Error("there sould be a single buy on a market")
	}
	if len(pool.Sells) != 0 {
		t.Error("There should not be any active sells on a market")
	}

}
