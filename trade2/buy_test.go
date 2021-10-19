package trade2

import (
	"reflect"
	"sync_tree/calc"
	"testing"
)

func TestInsertBuy(t *testing.T) {
	newPool := CreatePool()
	badTrade := CreateTrade(calc.Rand(), 10, 100)
	midTrade := CreateTrade(calc.Rand(), 50, 50)
	goodTrade := CreateTrade(calc.Rand(), 100, 10)
	newPool.insertBuy(badTrade)
	newPool.insertBuy(goodTrade)
	newPool.insertBuy(midTrade)
	if len(newPool.Buys) != 3 {
		t.Error("pool buys length should be equal to 3")
		return
	}
	if !reflect.DeepEqual(&newPool.Buys[0], &goodTrade) {
		t.Error("wrong start trade")
	}
	if !reflect.DeepEqual(&newPool.Buys[1], &midTrade) {
		t.Error("wrong mid trade")
	}
	if !reflect.DeepEqual(&newPool.Buys[2], &badTrade) {
		t.Error("wrong end trade")
	}
}

func TestEjectFirstBuy(t *testing.T) {
	newPool := CreatePool()
	badTrade := CreateTrade(calc.Rand(), 10, 100)
	goodTrade := CreateTrade(calc.Rand(), 100, 10)
	newPool.insertBuy(goodTrade)
	newPool.insertBuy(badTrade)
	ejected := newPool.ejectFirstBuy()
	if len(newPool.Buys) != 1 {
		t.Error("new length of buys should be equal to one")
	}
	if !reflect.DeepEqual(&ejected, &goodTrade) {
		t.Error("trade was ejected incorrectly,", ejected)
	}
}

func TestOperateBuyInEmptyPool(t *testing.T) {
	newPool := CreatePool()
	buy := CreateTrade(calc.Rand(), 10, 10)
	newPool.OperateBuy(buy)
	if len(newPool.Buys) != 1 {
		t.Error("Trade should be added")
	}
	if !reflect.DeepEqual(&newPool.Buys[0], &buy) {
		t.Error("Buy was added incorrectly")
	}
}
