package market

import (
	"orb/calc"
	"orb/database"
	"testing"
)

func TestMarketSave(t *testing.T) {
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
	mkt := Get(adress)
	mkt.Descr = "some new stuff"
	mkt.Save()
	lookedMkt := Look(adress)
	if lookedMkt.Descr != "some new stuff" {
		t.Error("Market have been saved but info have not been updated")
	}
	database.TestRM([]byte(dummyName))
	database.TestRM(adress)
}
