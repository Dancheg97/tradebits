package market

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestOperate(t *testing.T) {
	for i := 0; i < 100; i++ {
		firstRandTrade := Trade{
			Adress: []byte("a"),
		}
		secondRandTrade := Trade{
			Adress: []byte("b"),
		}
		randNumbers := []uint64{}
		for i := 0; i < 4; i++ {
			min := 1
			max := 100
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
				return
			}
			if len(outputs) != 0 {
				t.Error("if trades dont match, there should be 0 outputs")
				return
			}
			if !reflect.DeepEqual(trades[0], firstRandTrade) {
				t.Error("if trades dont match, first output trade should be the same to input")
				return
			}
			if !reflect.DeepEqual(trades[1], secondRandTrade) {
				t.Error("if trades dont match, second output trade should be the same to input")
				return
			}
		} else {
			if len(trades) != 1 {
				t.Error("if trades match, output should be one trade")
				return
			}
			if len(outputs) != 2 {
				t.Error("if trades match, there should be 2 outputs")
				return
			}
			sumInpMarket := uint64(0)
			sumInpMain := uint64(0)
			if firstRandTrade.IsSell {
				sumInpMarket = sumInpMarket + firstRandTrade.Offer
				sumInpMain = sumInpMain + secondRandTrade.Offer
			} else {
				sumInpMain = sumInpMain + firstRandTrade.Offer
				sumInpMarket = sumInpMarket + secondRandTrade.Offer
			}
			sumOutputMarket := uint64(0)
			sumOutputMain := uint64(0)
			for _, out := range outputs {
				sumOutputMarket = sumOutputMarket + out.MarketOut
				sumOutputMain = sumOutputMain + out.MainOut
			}
			if trades[0].IsSell {
				sumOutputMarket = sumOutputMarket + trades[0].Offer
			} else {
				sumOutputMain = sumOutputMain + trades[0].Offer
			}
			if sumInpMarket != sumOutputMarket {
				t.Error("input and output for market should be the same")
				return
			}
			if sumInpMain != sumOutputMain {
				t.Error("input and output for main should be the same")
				return
			}
		}
	}
}
