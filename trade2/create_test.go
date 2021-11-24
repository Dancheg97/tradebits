package trade

import (
	"reflect"
	"testing"
)

func TestCreateTrade(t *testing.T) {
	trd := CreateTrade([]byte{0}, 1, 2)
	if !reflect.DeepEqual(trd.Adress, []byte{0}) {
		t.Error("wrong adress has been passed")
	}
	if trd.Offer != 1 {
		t.Error("wrong offer has been passed")
	}
	if trd.Recieve != 2 {
		t.Error("wrong recieve amount has been passed")
	}
}

func TestCreatePool(t *testing.T) {
	pool := CreatePool()
	if pool.Buys == nil || pool.Sells == nil {
		t.Error("buys&sells field should be initialized")
	}
	if len(pool.Buys) != 0 || len(pool.Sells) != 0 {
		t.Error("len of buys and sells should be equal to zero")
	}
}
