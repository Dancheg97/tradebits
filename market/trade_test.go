package market

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestOperate(t *testing.T) {
	for i := 0; i < 10000; i++ {
		firstRandTrade := Trade{
			Adress: []byte("a"),
		}
		secondRandTrade := Trade{
			Adress: []byte("b"),
		}
		randNumbers := []uint64{}
		for i := 0; i < 4; i++ {
			min := 0
			max := 3000000000
			randNum := rand.Intn(max-min) + min
			randNumbers = append(randNumbers, uint64(randNum))
		}
		firstRandTrade.Offer = randNumbers[0]
		firstRandTrade.Recieve = randNumbers[1]
		secondRandTrade.Offer = randNumbers[2]
		secondRandTrade.Recieve = randNumbers[3]
		randBool := rand.Intn(2) != 0
		firstRandTrade.IsSell = randBool
		secondRandTrade.IsSell = !randBool
		match, trades, outputs := firstRandTrade.operate(secondRandTrade)
		if !match {
			if len(trades) != 2 {
				t.Error("if trades dont match, there should 2 trades output")
			}
			if len(outputs) != 0 {
				t.Error("if trades dont match, there should be 0 outputs")
			}
			if !reflect.DeepEqual(trades[0], firstRandTrade) {
				t.Error("if trades dont match, first output trade should be the same to input")
			}
			if !reflect.DeepEqual(trades[1], secondRandTrade) {
				t.Error("if trades dont match, second output trade should be the same to input")
			}
		} else {
			if len(trades) != 1 {
				t.Error("if trades match, output should be one trade")
			}
			if len(outputs) != 2 {
				t.Error("if trades match, there should be 2 outputs")
			}
			inpMarket := uint64(0)
			inpMain := uint64(0)
			
		}
	}
}
