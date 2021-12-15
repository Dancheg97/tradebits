package market

import (
	"orb/calc"
	"orb/database"
	"reflect"
	"testing"
)

func TestMarketLook(t *testing.T) {
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
	mkt := Look(adress)
	if !reflect.DeepEqual(mkt.MesKey, dummyMessageKey) {
		t.Error("keys are not the same, look market error")
	}
	database.TestRM([]byte(dummyName))
	database.TestRM(adress)
}
