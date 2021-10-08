package market

import (
	"reflect"
	"sync_tree/calc"
	"sync_tree/data"
	"testing"
)

func TestGetFreeMarket(t *testing.T) {
	var adress = calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	Create(
		adress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	market := Get(adress)
	if !reflect.DeepEqual(market.Name, dummyName) {
		t.Error("keys are not the same, get asset error")
	}
	market.Save()
	data.TestRM([]byte(dummyName))
	data.TestRM(adress)
}

func TestGetMarketThatDontExist(t *testing.T) {
	var adress = calc.Rand()
	mkt := Get(adress)
	if mkt != nil {
		t.Error("that test should not return anything")
	}
}
