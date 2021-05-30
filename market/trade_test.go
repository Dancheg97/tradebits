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
			num := uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
			randNumbers = append(randNumbers, num)
		}
		sell.Offer = randNumbers[0]
		sell.Recieve = randNumbers[1]
		buy.Offer = randNumbers[2]
		buy.Recieve = randNumbers[3]
		sell.operate(buy)
		
	}
}
