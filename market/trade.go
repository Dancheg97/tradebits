package market

// import "sync_tree/user"

// type Buy struct {
// 	Adress  []byte
// 	Offer   uint64
// 	Recieve uint64
// }

// type Sell struct {
// 	Adress  []byte
// 	Offer   uint64
// 	Recieve uint64
// }

// func (b *Buy) match(s *Sell) (bool, bool) {
// 	if b.Offer == s.Offer && b.Recieve == s.Recieve {
// 		user.Get()
// 	}
// }

// func (s Sell) match(b Buy) {

// }

// func (m market) addBuy(b Buy) {

// }

// func (m market) addSell(s Sell) {

// }

// func (m market) cancelBuy(adress []byte) {

// }

// func (m market) cancelSell(adress []byte) {

// }

// // type Trade struct {
// // 	Adress  []byte
// // 	IsSell  bool
// // 	Offer   uint64
// // 	Recieve uint64
// // }

// // /*
// // Recursive function to add trades to existing market. Each new iteration
// // */
// // func (m *market) OperateTrade(newTrade Trade) bool {
// // 	m.OpCount = m.OpCount + 1
// // 	if newTrade.Offer == 0 || newTrade.Recieve == 0 {
// // 		return false
// // 	}
// // 	if newTrade.IsSell {
// // 		if len(m.Buys) == 0 {
// // 			m.Sells = append(m.Sells, newTrade)
// // 			return true
// // 		}
// // 		trades, outputs := newTrade.operate(m.Buys[0])
// // 		m.Buys = m.Buys[1:]
// // 		m.outputs = append(m.outputs, outputs...)
// // 		m.output()
// // 		if len(trades) == 2 {
// // 			m.addTrade(trades[0])
// // 			m.addTrade(trades[1])
// // 			return true
// // 		}
// // 		if len(trades) == 1 {
// // 			m.OperateTrade(newTrade)
// // 		}
// // 		return true
// // 	} else {
// // 		if len(m.Sells) == 0 {
// // 			m.Buys = append(m.Buys, newTrade)
// // 			return true
// // 		}
// // 		trades, outputs := newTrade.operate(m.Sells[0])
// // 		m.Sells = m.Sells[1:]
// // 		m.outputs = append(m.outputs, outputs...)
// // 		m.output()
// // 		if len(trades) == 2 {
// // 			m.addTrade(trades[0])
// // 			m.addTrade(trades[1])
// // 			return true
// // 		}
// // 		if len(trades) == 1 {
// // 			m.OperateTrade(newTrade)
// // 		}
// // 		return true
// // 	}
// // }

// // // function that returns offers from some adress to trade creator
// // func (m *market) CancelTrades(adress []byte) bool {
// // 	for index, buyTrade := range m.Buys {
// // 		if reflect.DeepEqual(adress, buyTrade.Adress) {
// // 			m.Buys = append(m.Buys[:index], m.Buys[index+1:]...)
// // 			u := user.Get(adress)
// // 			if u == nil {
// // 				return false
// // 			}
// // 			u.Balance = u.Balance + buyTrade.Offer
// // 			u.Save()
// // 		}
// // 	}
// // 	for index, sellTrade := range m.Sells {
// // 		if reflect.DeepEqual(adress, sellTrade.Adress) {
// // 			m.Sells = append(m.Sells[:index], m.Sells[index+1:]...)
// // 			u := user.Get(adress)
// // 			if u == nil {
// // 				return false
// // 			}
// // 			u.Markets[string(m.adress)] = u.Markets[string(m.adress)] + sellTrade.Offer
// // 			u.Save()
// // 		}
// // 	}
// // 	return true
// // }

// // // function to send all outputs back to users
// // func (m *market) output() {
// // 	for _, output := range m.outputs {
// // 		for {
// // 			u := user.Get(output.Adress)
// // 			if u != nil {
// // 				u.Balance = u.Balance + output.MainOut
// // 				marketAdr := string(m.adress)
// // 				u.Markets[marketAdr] = u.Markets[marketAdr] + output.MarketOut
// // 				u.Save()
// // 				break
// // 			}
// // 			time.Sleep(time.Second)
// // 		}
// // 	}
// // }

// // func (new Trade) operate(old Trade) ([]Trade, []output) {
// // 	if new.Recieve < old.Offer {
// // 		ratio := float64(old.Recieve) / float64(old.Offer)
// // 		potentialNewOffer := uint64(math.Ceil(float64(new.Recieve) * ratio))
// // 		if potentialNewOffer > new.Offer {
// // 			return []Trade{new, old}, nil
// // 		}
// // 		newOutput := output{Adress: new.Adress}
// // 		oldOutput := output{Adress: old.Adress}
// // 		if old.IsSell {
// // 			newOutput.MainOut = new.Offer - potentialNewOffer
// // 			newOutput.MarketOut = new.Recieve
// // 			oldOutput.MainOut = potentialNewOffer
// // 		} else {
// // 			newOutput.MarketOut = new.Offer - potentialNewOffer
// // 			newOutput.MainOut = new.Recieve
// // 			oldOutput.MarketOut = potentialNewOffer
// // 		}
// // 		old.Offer = old.Offer - new.Recieve // add dot after old in var name
// // 		old.Recieve = old.Recieve - potentialNewOffer
// // 		return []Trade{old}, []output{newOutput, oldOutput}
// // 	}
// // 	newRatio := float64(new.Recieve) / float64(new.Offer)
// // 	oldRatio := float64(old.Offer) / float64(old.Recieve)
// // 	if newRatio > oldRatio {
// // 		return []Trade{new, old}, nil
// // 	}
// // 	newOutput := output{Adress: new.Adress}
// // 	oldOutput := output{Adress: old.Adress}
// // 	if new.IsSell {
// // 		newOutput.MainOut = old.Offer
// // 		oldOutput.MarketOut = old.Recieve
// // 	} else {
// // 		newOutput.MarketOut = old.Offer
// // 		oldOutput.MainOut = old.Recieve
// // 	}
// // 	new.Offer = new.Offer - old.Recieve
// // 	new.Recieve = new.Recieve - old.Offer
// // 	if new.Offer == 0 && new.Recieve == 0 {
// // 		return []Trade{}, []output{newOutput, oldOutput}
// // 	}
// // 	return []Trade{new}, []output{newOutput, oldOutput}
// // }

// // // assistive func to add trade to curr trade list in proper place by ratio
// // func (m *market) addTrade(t Trade) {
// // 	currRatio := float64(t.Offer) / float64(t.Recieve)
// // 	if t.IsSell {
// // 		if len(m.Sells) == 0 {
// // 			m.Sells = append(m.Sells, t)
// // 			return
// // 		}
// // 		for index, sell := range m.Sells {
// // 			sellRatio := float64(sell.Offer) / float64(sell.Recieve)
// // 			if currRatio > sellRatio {
// // 				m.Sells = append(m.Sells[:index+1], m.Sells[index:]...)
// // 				m.Sells[index] = t
// // 			}
// // 		}
// // 	} else {
// // 		if len(m.Buys) == 0 {
// // 			m.Buys = append(m.Buys, t)
// // 			return
// // 		}
// // 		for index, buy := range m.Buys {
// // 			buyRatio := float64(buy.Offer) / float64(buy.Recieve)
// // 			if currRatio > buyRatio {
// // 				m.Buys = append(m.Buys[:index+1], m.Buys[index:]...)
// // 				m.Buys[index] = t
// // 			}
// // 		}
// // 	}
// // }
