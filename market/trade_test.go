package market

import (
	"math/rand"
	"testing"
	"time"
	//"time"
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
		firstRandTrade.operate(secondRandTrade)
		
	}
}
