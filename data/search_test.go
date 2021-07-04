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
	if reflect.DeepEqual(rez[0], adr) {
		return
	}
	t.Error("adresses should not match")
}

func TestSearchChange(t *testing.T) {
	name := "name2"
	fistAdress := []byte{1, 1, 1}
	secondAdress := []byte{1, 1, 2}
	SearchAdd(name, fistAdress)
	SearchChange(name, secondAdress)
	rez := Search(name)
	if reflect.DeepEqual(rez[0], name) {
		
	}
}
