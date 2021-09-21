package search

import (
	"os"
	"path"
	"reflect"
	"runtime"
	"sync_tree/calc"
	"testing"
	"time"
)

func TestSearchAdd(t *testing.T) {
	name := "name22"
	adr := calc.Rand()
	Add(name, adr)
	searcher.Delete(string(adr))
}

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

func TestSearchChange(t *testing.T) {
	adr := calc.Rand()
	name1 := "Xname22"
	name2 := "Fname22"
	Add(name1, adr)
	Change(name2, adr)
	rez1 := Search(name2)
	if !reflect.DeepEqual(rez1[0], adr) {
		t.Error("found name should be equal to second")
	}
	searcher.Delete(string(adr))
}

func TestSearchAddDifferentAdressesToSameName(t *testing.T) {
	firstAdress := calc.Rand()
	secondAdress := calc.Rand()
	Add("name", firstAdress)
	Add("name", secondAdress)
	search := Search("name")
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

func TestRecreateSearcher(t *testing.T) {
	time.Sleep(time.Second * 8)
	searcher.Close()
	_, filename, _, _ := runtime.Caller(0)
	searchPath := path.Dir(filename) + "/bleve"
	os.RemoveAll(searchPath)
	openSearch()
}

func TestSearchWithAttemptOfNameSubstitution(t *testing.T) {
	firstAdress := calc.Rand()
	secondAdress := calc.Rand()
	Add("name", firstAdress)
	Add("name", secondAdress)
	Change("another", secondAdress)
	firstSearch := Search("name")
	secondSearch := Search("another")
	if !reflect.DeepEqual(firstSearch[0], firstAdress) {
		t.Error("first adress not matching")
	}
	if !reflect.DeepEqual(secondSearch[0], secondAdress) {
		t.Error("second adress not matching")
	}
	searcher.Delete(string(firstAdress))
	searcher.Delete(string(secondAdress))
}
