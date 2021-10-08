package market

import (
	"sync_tree/calc"
	"sync_tree/data"
	"sync_tree/trade"
	"sync_tree/user"
	"testing"
)

var dummyMessageKey = []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2}
var dummyImageLink = "test.imagelink/thereisnoimagebythislink"
var dummyDescription = `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrbled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets ctaining Lorem Ipsum pages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`
var dummyInputFee = uint64(100)
var dummyOutputFee = uint64(100)
var dummyWorkTime = "+3GMT 9:00 - 21:00"
var dummyDelimiter = uint64(2)

func TestOperateOutput(t *testing.T) {
	dummyName := string(calc.Rand()[0:8])
	dummyUserAdress := calc.Rand()
	dummyMarketAdress := calc.Rand()
	user.Create(
		dummyUserAdress,
		dummyMessageKey,
		dummyName,
	)
	output := trade.Output{
		Adress: dummyUserAdress,
		Main:   100,
		Market: 50,
	}
	operateOutput(output, dummyMarketAdress)
	u := user.Get(dummyUserAdress)
	if u.Balance != 100 || u.Balances[string(dummyMarketAdress)] != 50 {
		t.Error("output was operated incorrectly")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(dummyUserAdress)
}

