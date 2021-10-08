package market

import (
	"sync_tree/trade"
	"sync_tree/user"
)

// This function is operating output for some trade and market adress
func operateOutput(t trade.Output, marketAdress []byte) {
	u := user.Get(t.Adress)
	u.Balance = u.Balance + t.Main
	u.Balances[string(marketAdress)] = u.Balances[string(marketAdress)] + t.Market
	u.Save()
}

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
