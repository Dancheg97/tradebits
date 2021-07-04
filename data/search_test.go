package data

import (
	"reflect"
	"testing"
)

func TestSearchAdd(t *testing.T) {
	name := "name"
	adr := []byte{1, 2, 3}
	SearchAdd(name, adr)
	searcher.Delete(name)
}

func TestSearchSearch(t *testing.T) {
	name := "name1"
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
	name1 := "Xname"
	name2 := "Fname"
	SearchAdd(name1, adr)
	SearchChange(name2, adr)
	rez := Search(name2)
	if reflect.DeepEqual(rez[0], name2) {
		return
	}
	t.Error("found name should be equal to second")
}
