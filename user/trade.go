package user

import (
	"sync_tree/trade"
)

// This function is bounding specific sell function to user, if its not
// possible, returns false, if operated successully, returns true
func (u *user) AttachSell(sell *trade.Sell, marketAdress []byte) bool {
	if u.adress == nil {
		return false
	}
	if sell.Adress != nil {
		return false
	}
	if sell.Offer == 0 || sell.Recieve == 0 {
		return false
	}
	if val, ok := u.Balances[string(marketAdress)]; ok {
		if val >= sell.Offer {
			u.Balances[string(marketAdress)] = val - sell.Offer
			sell.Adress = u.adress
			return true
		}
	}
	return false
}

// this function is bounding specific buy to user, if its not possible returns
// false, if operated successfully returns true
func (u *user) AttachBuy(buy *trade.Buy) bool {
	if u.adress == nil {
		return false
	}
	if buy.Adress != nil {
		return false
	}
	if buy.Offer == 0 || buy.Recieve == 0 {
		return false
	}
	if u.Balance < buy.Offer {
		return false
	}
	u.Balance = u.Balance - buy.Offer
	buy.Adress = u.adress
	return true
}
