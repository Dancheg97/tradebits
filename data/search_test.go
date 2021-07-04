package data

import (
	"reflect"
	"testing"
)

func TestSearchAdd(t *testing.T) {
	name := "name22"
	adr := []byte{1, 2, 3}
	SearchAdd(name, adr)
	searcher.Delete(name)
}

func TestSearchSearch(t *testing.T) {
	name := "name122"
	adr := []byte{1, 2, 3}
	SearchAdd(name, adr)
	rez := Search(name)
	if !reflect.DeepEqual(rez[0], adr) {
		t.Error("adresses should not match")
	}
	searcher.Delete(name)
}

func TestSearchChange(t *testing.T) {
	adr := []byte{1, 1, 1}
	name1 := "Xname22"
	name2 := "Fname22"
	SearchAdd(name1, adr)
	SearchChange(name2, adr)
	rez := Search(name2)
	if !reflect.DeepEqual(rez[0], adr) {
		t.Error("found name should be equal to second")
	}
	searcher.Delete("name2")
}

func TestSearchAddMultipleAdressesOnSameName(t *testing.T) {

}
