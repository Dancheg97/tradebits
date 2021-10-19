package trade2

import (
	"reflect"
	"sync_tree/calc"
	"testing"
)

// This file contains comlpex tests for trade logic

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

func TestOperateMatchingBuyAndSell(t *testing.T) {
	newPool := CreatePool()
	sell := CreateTrade(calc.Rand(), 10, 10)
	buy := CreateTrade(calc.Rand(), 10, 10)
	newPool.OperateSell(sell)
	newPool.OperateBuy(buy)
	if len(newPool.Buys) != 0 || len(newPool.Sells) != 0 {
		t.Error("threre should not be any trades")
	}
	if len(newPool.MainOutputs) != 1 || len(newPool.MarketOutputs) != 1 {
		t.Error("threre should be a single output for market and main")
	}
}
