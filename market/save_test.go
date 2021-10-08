package market

import (
	"sync_tree/calc"
	"testing"
)

func TestMarketSave(t *testing.T) {
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
	mkt := Get(adress)
	mkt.Descr = "some new stuff"
	mkt.Save()
	lookedMkt := Look(adress)
	if lookedMkt.Descr != "some new stuff" {
		t.Error("Market have been saved but info have not been updated")
	}
}
