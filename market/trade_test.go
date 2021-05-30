package market

import "testing"

func TestOperate(t *testing.T) {
	buy := Trade{
		Adress:  []byte("b"),
		IsSell:  true,
		Offer:   400,
		Recieve: 300,
	}
	sell := Trade{
		Adress:  []byte("ss"),
		IsSell:  false,
		Offer:   400,
		Recieve: 600,
	}
	operated, trades, outputs := buy.operate(sell)
	t.Error(operated)
	t.Error(trades)
	t.Error(outputs)
}
