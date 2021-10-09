package market

import (
	"sync_tree/calc"
	"sync_tree/data"
	"testing"
)

func TestHasNoTrades(t *testing.T) {
	dummyAdress := calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	Create(
		dummyAdress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	mkt := Get(dummyAdress)
	dummyUserAdress := calc.Rand()
	hasTrades := mkt.HasTrades(dummyUserAdress)
	if hasTrades {
		t.Error("new market should not have any trades")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(dummyAdress)
}
