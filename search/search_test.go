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
	adr := []byte{1, 2, 3}
	Add(name, adr)
	searcher.Delete(string(adr))
}

func TestSearchSearch(t *testing.T) {
	name := "name122"
	adr := []byte{1, 2, 3}
	Add(name, adr)
	rez := Search(name)
	if !reflect.DeepEqual(rez[0], adr) {
		t.Error("adresses should not match")
	}
	searcher.Delete(string(adr))
}

func TestSearchChange(t *testing.T) {
	adr := []byte{1, 1, 1}
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

// func TestSearchChange1name2adresses(t *testing.T) {
// 	firstAdress := []byte{1, 2, 3}
// 	secondAdress := []byte{1, 2, 3}
// 	Add("name", firstAdress)
// 	Add("name", secondAdress)
// 	Change("anotherName", secondAdress)
// 	firstSearch := Search("name")
// 	secondSearch := Search("antoherName")
// 	if !reflect.DeepEqual(firstSearch[0], firstAdress) {
// 		t.Error("first adress not matching")
// 	}
// 	if !reflect.DeepEqual(secondSearch[0], secondAdress) {
// 		t.Error("second adress not matching")
// 	}
// }

// func TestSearchAddDifferentAdressesToSameName(t *testing.T) {
// 	firstAdress := []byte{1, 2, 3}
// 	secondAdress := []byte{1, 2, 3, 4}
// 	Add("name", firstAdress)
// 	Add("name", secondAdress)
// 	firstSearch := Search("name")
// 	secondSearch := Search("name")
// 	if !reflect.DeepEqual(firstSearch[0], firstAdress) {
// 		t.Error("first adress not matching")
// 	}
// 	if !reflect.DeepEqual(secondSearch[0], secondAdress) {
// 		t.Error("second adress not matching")
// 	}
// }

func TestSearchAddMultipleAdressesOnSameName(t *testing.T) {
	adr1 := []byte{0, 1, 2, 3}
	adr2 := []byte{0, 1, 2, 4}
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
	time.Sleep(time.Second * 3)
	searcher.Close()
	_, filename, _, _ := runtime.Caller(0)
	searchPath := path.Dir(filename) + "/bleve"
	os.RemoveAll(searchPath)
	openSearch()
}
