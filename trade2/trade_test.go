package trade2

import (
	"sync_tree/calc"
	"testing"
)

func TestCreateTrade(t *testing.T) {
	trd := CreateTrade(calc.Rand(), 100, 120)
	if trd == nil {
		t.Error("This trade should be created normally")
	}
}

func TestCreateTradeError(t *testing.T) {
	trd := CreateTrade(calc.Rand(), 0, 12)
	if trd != nil {
		t.Error("this trade cannot be created")
	}
}
