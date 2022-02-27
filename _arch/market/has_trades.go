package market

import "reflect"

// making a check, wether some user has trades on that market
func (m *market) HasTrades(adress []byte) bool {
	for _, trade := range m.Pool.Buys {
		if reflect.DeepEqual(trade.Adress, adress) {
			return true
		}
	}
	for _, trade := range m.Pool.Sells {
		if reflect.DeepEqual(trade.Adress, adress) {
			return true
		}
	}
	return false
}
