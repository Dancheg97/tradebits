package market

import (
	"math/rand"
	"testing"
)

func TestOperate(t *testing.T) {
	for i := 0; i < 10000; i++ {
		sell := Trade{
			Adress: []byte("b"),
			IsSell: true,
		}
		buy := Trade{
			Adress: []byte("c"),
			IsSell: false,
		}
		randNumbers := []uint64{}
		for i := 0; i < 4; i++ {
			min := 0
			max := 30000000
			randNum := rand.Intn(max-min) + min
			randNumbers = append(randNumbers, uint64(randNum))
		}
		sell.Offer = randNumbers[0]
		sell.Recieve = randNumbers[1]
		buy.Offer = randNumbers[2]
		buy.Recieve = randNumbers[3]
		operated, trades, outputs := sell.operate(buy)
		if operated {
			if len(trades) != 1 {
				t.Error("there should be only one outgoing trade")
			}
			if 
		}
	}
}
