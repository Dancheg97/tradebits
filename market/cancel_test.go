package market

import (
	"sync_tree/calc"
	"sync_tree/data"
	"testing"
)

func TestCancelBuy(t *testing.T) {
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
	
	data.TestRM([]byte(dummyName))
	data.TestRM(dummyAdress)
}
