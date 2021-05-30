package market

import "testing"

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
		
	}
}
