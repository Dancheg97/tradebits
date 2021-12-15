package market

import (
	"orb/calc"
	"orb/database"
	"reflect"
	"testing"
)

func TestGetFreeMarket(t *testing.T) {
	adress := calc.Rand()
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
	database.TestRM([]byte(dummyName))
	database.TestRM(adress)
}

func TestGetMarketThatDontExist(t *testing.T) {
	adress := calc.Rand()
	mkt := Get(adress)
	if mkt != nil {
		t.Error("that test should not return anything")
	}
}
