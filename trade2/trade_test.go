package trade2

import (
	"reflect"
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
	firstTrade := CreateTrade(calc.Rand(), 7, 10)
	secondTrade := CreateTrade(calc.Rand(), 8, 6)
	firstOutput, secondOutput := firstTrade.close(secondTrade)
	if firstOutput != nil && secondOutput != nil {
		t.Error("Those trades should not be operated, ", firstOutput, secondOutput)
	}
}

func TestTradeFirstWillCloseSecond(t *testing.T) {
	firstAdress := calc.Rand()
	secondAdress := calc.Rand()
	firstTrade := CreateTrade(firstAdress, 7, 10)
	secondTrade := CreateTrade(secondAdress, 9, 5)
	firstNanOut, secondNanOut := secondTrade.close(firstTrade)
	if firstNanOut != nil || secondNanOut != nil {
		t.Error("second trade should be unable to return pointer value")
	}
	t.Error(firstNanOut, secondNanOut)
	firstOutput, secondOutput := firstTrade.close(secondTrade)
	if firstOutput == nil || secondOutput == nil {
		t.Error("Those trades should be operated")
	}
	if !reflect.DeepEqual(firstOutput.Adress, firstAdress) {
		t.Error("first adress is not matching")
	}
	if !reflect.DeepEqual(secondOutput.Adress, secondAdress) {
		t.Error("second adress is not matching")
	}
	if firstOutput.Amount != 9 {
		t.Error("first output should be 9")
	}
	if secondOutput.Amount != 5 {
		t.Error("second output should be 5")
	}
}
