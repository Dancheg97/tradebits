package market

import (
	"reflect"
	"sync_tree/trade"
	"sync_tree/user"
)

// keka
type market struct {
	adress    []byte
	Name      string
	MesKey    []byte
	Descr     string
	Img       string
	OpCount   uint64
	Pool      trade.TradePool
	InputFee  uint64
	OutputFee uint64
	WorkTime  string
	Delimiter uint64
	Users     [][]byte
}

// This function is operating output for some trade and market adress
func operateOutput(t trade.Output, marketAdress []byte) {
	u := user.Get(t.Adress)
	u.Balance = u.Balance + t.Main
	u.Balances[string(marketAdress)] = u.Balances[string(marketAdress)] + t.Market
	u.Save()
}

// attaches buy trade to market, you can't attach trade twice
func (m *market) AttachBuy(b *trade.Buy) bool {
	if m.adress == nil {
		return false
	}
	if b.Adress == nil {
		return false
	}
	m.Pool.OperateBuy(*b)
	for _, output := range m.Pool.Outputs {
		go operateOutput(output, m.adress)
	}
	m.Pool.Outputs = []trade.Output{}
	b.Adress = nil
	b.Offer = 0
	return true
}

// attaches sell trade to market, you can't attach trade twice
func (m *market) AttachSell(s *trade.Sell) bool {
	if m.adress == nil {
		return false
	}
	if s.Adress == nil {
		return false
	}
	m.Pool.OperateSell(*s)
	for _, output := range m.Pool.Outputs {
		go operateOutput(output, m.adress)
	}
	m.Pool.Outputs = []trade.Output{}
	s.Adress = nil
	s.Offer = 0
	return true
}

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

// this function cancelles trades
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
