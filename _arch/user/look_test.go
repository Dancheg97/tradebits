package user

import (
	"orb/calc"
	"orb/data"
	"reflect"
	"testing"
)

func TestUserLook(t *testing.T) {
	var adress = calc.Rand()
	dummyName := string(calc.Rand()[0:8])
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Look(adress)
	if len(usr.adress) != 0 {
		t.Error("user adress should be empty")
	}
	if reflect.DeepEqual(adress, usr.adress) {
		t.Error("user info is incorrect")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(adress)
}
