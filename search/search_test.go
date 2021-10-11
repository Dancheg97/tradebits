package search

import (
	"reflect"
	"sync_tree/calc"
	"testing"
	"time"
)

func TestSearchSearch(t *testing.T) {
	name := "name122"
	adr := calc.Rand()
	Add(name, adr)
	rez := Search(name)
	if !reflect.DeepEqual(rez[0], adr) {
		t.Error("adresses should not match")
	}
	searcher.Delete(string(adr))
}

func TestSearchAddDifferentAdressesToSameName(t *testing.T) {
	firstAdress := calc.Rand()
	secondAdress := calc.Rand()
	Add("namel", firstAdress)
	Add("namel", secondAdress)
	time.Sleep(time.Second * 2)
	search := Search("namel")
	firstCondition := (reflect.DeepEqual(search[0], firstAdress) &&
		reflect.DeepEqual(search[1], secondAdress))
	secondCondition := (reflect.DeepEqual(search[0], secondAdress) &&
		reflect.DeepEqual(search[1], firstAdress))
	if firstCondition || secondCondition {
		return
	}
	t.Error("conditions not satisfied", firstCondition, secondCondition)
}

func TestSearchAddMultipleAdressesOnSameName(t *testing.T) {
	adr1 := calc.Rand()
	adr2 := calc.Rand()
	sameName := "snm"
	Add(sameName, adr1)
	Add(sameName, adr2)
	rez := Search(sameName)
	if len(rez) != 2 {
		t.Error()
	}
	searcher.Delete(string(adr1))
	searcher.Delete(string(adr2))
}

func TestSearchOver30requests(t *testing.T) {
	adresses := [][]byte{}
	for i := 0; i < 35; i++ {
		adr := calc.Rand()
		Add("stuff ", adr)
		adresses = append(adresses, adr)
	}
	rez := Search("stuff")
	if len(rez) != 30 {
		t.Error("lenght: =>", len(rez))
	}
	for _, adr := range adresses {
		searcher.Delete(string(adr))
	}
}
