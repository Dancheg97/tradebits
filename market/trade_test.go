package market

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestOperate(t *testing.T) {
	for i := 0; i < 100; i++ {
		firstRandTrade := Trade{Adress: []byte("a")}
		secondRandTrade := Trade{Adress: []byte("b")}
		randNumbers := []uint64{}
		for i := 0; i < 4; i++ {
			rand.Seed(time.Now().UnixNano())
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
		trades, outputs := firstRandTrade.operate(secondRandTrade)
		if outputs == nil {
			if !reflect.DeepEqual(firstRandTrade, trades[0]) {
				t.Error("if trades dont operate they should be same")
			}
			if !reflect.DeepEqual(secondRandTrade, trades[1]) {
				t.Error("if trades dont operate they should be the smae")
			}
		} else {
			
		}
	}
}
