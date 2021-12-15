package market

import (
	"orb/user"
	"reflect"
)

func cancelBuy(userAdress []byte, offer uint64) {
	usr := user.Get(userAdress)
	usr.Balance = usr.Balance + offer
	usr.Save()
}

func cancelSell(userAdress []byte, marketAdress []byte, offer uint64) {
	usr := user.Get(userAdress)
	mktAdress := string(marketAdress)
	usr.Balances[mktAdress] = usr.Balances[mktAdress] + offer
	usr.Save()
}

func (m *market) CancelTrades(adress []byte) {
	for idx, trade := range m.Pool.Buys {
		if reflect.DeepEqual(trade.Adress, adress) {
			go cancelBuy(adress, trade.Offer)
			m.Pool.Buys = append(m.Pool.Buys[:idx], m.Pool.Buys[idx+1:]...)
		}
	}
	for idx, trade := range m.Pool.Sells {
		if reflect.DeepEqual(trade.Adress, adress) {
			go cancelSell(adress, m.adress, trade.Offer)
			m.Pool.Sells = append(m.Pool.Sells[:idx], m.Pool.Sells[idx+1:]...)
		}
	}
}
