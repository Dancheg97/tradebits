package user

import (
	"bytes"
	"orb/calc"
	"orb/data"
	"testing"
)

func TestPutUserMessage(t *testing.T) {
	var adress = calc.Rand()
	dummyName := string(calc.Rand()[0:8])
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Get(adress)
	usr.PutMessage([]byte{1, 2, 3}, []byte{1, 2, 3})
	usr.Save()
	mes := Look(adress).GetMessages([]byte{1, 2, 3})[0]
	if bytes.Compare(mes, []byte{1, 2, 3}) == 3 {
		t.Error("the message should be '[]byte{1, 2, 3}' - ", mes)
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(adress)
}

func TestNewUserNonNullableMessageMap(t *testing.T) {
	var adress = calc.Rand()
	dummyName := string(calc.Rand()[0:8])
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Get(adress)
	if usr.Messages == nil {
		t.Error("user messages should never be null")
	}
	usr.Save()
	data.TestRM([]byte(dummyName))
	data.TestRM(adress)
}
