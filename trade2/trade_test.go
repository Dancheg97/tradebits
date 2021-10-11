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

func TestTradeFirstWontCloseSecond(t *testing.T) {
	firstTrade := CreateTrade(calc.Rand(), 119, 181)
	secondTrade := CreateTrade(calc.Rand(), 120, 182)
	firstOutput, secondOutput := firstTrade.close(secondTrade)
	if firstOutput != nil && secondOutput != nil {
		t.Error("coca cola", firstOutput, secondOutput)
	}
}

