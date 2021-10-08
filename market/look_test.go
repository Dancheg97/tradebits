package market

import (
	"reflect"
	"sync_tree/calc"
	"sync_tree/data"
	"testing"
)

func TestMarketLook(t *testing.T) {
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
	mkt := Look(adress)
	if !reflect.DeepEqual(mkt.MesKey, dummyMessageKey) {
		t.Error("keys are not the same, look asset error")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(adress)
}
