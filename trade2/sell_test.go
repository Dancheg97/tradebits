package trade2

import (
	"reflect"
	"sync_tree/calc"
	"testing"
)

func TestInsertSell(t *testing.T) {
	newPool := CreatePool()
	badTrade := CreateTrade(calc.Rand(), 10, 100)
	midTrade := CreateTrade(calc.Rand(), 50, 50)
	goodTrade := CreateTrade(calc.Rand(), 100, 10)
	newPool.insertSell(badTrade)
	newPool.insertSell(goodTrade)
	newPool.insertSell(midTrade)
	if len(newPool.Sells) != 3 {
		t.Error("pool buys length should be equal to 3")
		return
	}
	if !reflect.DeepEqual(&newPool.Sells[0], &goodTrade) {
		t.Error("wrong start trade")
	}
	if !reflect.DeepEqual(&newPool.Sells[1], &midTrade) {
		t.Error("wrong mid trade")
	}
	if !reflect.DeepEqual(&newPool.Sells[2], &badTrade) {
		t.Error("wrong end trade")
	}
}

func TestEjectFirstSell(t *testing.T) {
	newPool := CreatePool()
	badTrade := CreateTrade(calc.Rand(), 10, 100)
	goodTrade := CreateTrade(calc.Rand(), 100, 10)
	newPool.insertSell(goodTrade)
	newPool.insertSell(badTrade)
	ejected := newPool.ejectFirstSell()
	if len(newPool.Sells) != 1 {
		t.Error("new length of buys should be equal to one")
	}
	if !reflect.DeepEqual(&ejected, &goodTrade) {
		t.Error("trade was ejected incorrectly,", ejected)
	}
}

func TestOperateSellInEmptyPool(t *testing.T) {
	newPool := CreatePool()
	sell := CreateTrade(calc.Rand(), 10, 10)
	newPool.OperateSell(sell)
	if len(newPool.Sells) != 1 {
		t.Error("Trade should be added")
	}
	if !reflect.DeepEqual(&newPool.Sells[0], &sell) {
		t.Error("Sell was added incorrectly")
	}
}
