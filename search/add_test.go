package search

import (
	"orb/calc"
	"testing"
)

func TestAdd(t *testing.T) {
	name := "name22"
	adr := calc.Rand()
	Add(name, adr)
	searcher.Delete(string(adr))
}
