package trade

// this struct is used only to transfer data about market outputs for some user
type Output struct {
	Adress []byte
	Main   uint64
	Market uint64
}

type TradePool struct {
	Buys    []Buy
	Sells   []Sell
	Outputs []Output
}

// all trades are alwayts closing to the side better side
func match(
	buy *Buy,
	sell *Sell,
) []Output {
	if float64(buy.Offer)/float64(sell.Recieve) >= float64(buy.Recieve/sell.Offer) {
		if buy.Offer < sell.Recieve && buy.Recieve < sell.Offer {
			defer buy.close()
			defer sell.reduceOffer(buy.Recieve)
			defer sell.reduceRecieve(buy.Offer)
			buyerOutput := Output{
				Adress: buy.Adress,
				Market: buy.Recieve,
			}
			sellerOutput := Output{
				Adress: sell.Adress,
				Main:   buy.Offer,
			}
			return []Output{
				buyerOutput,
				sellerOutput,
			}
		}
		if sell.Offer < buy.Recieve && sell.Recieve < buy.Offer {
			defer sell.close()
			defer buy.reduceOffer(sell.Recieve)
			defer buy.reduceRecieve(sell.Offer)
			buyerOutput := Output{
				Adress: buy.Adress,
				Market: sell.Offer,
			}
			sellerOutput := Output{
				Adress: sell.Adress,
				Main:   sell.Recieve,
			}
			return []Output{
				buyerOutput,
				sellerOutput,
			}
		}
		if sell.Offer >= buy.Recieve && buy.Offer >= sell.Recieve {
			defer sell.close()
			defer buy.close()
			buyerOutput := Output{
				Adress: buy.Adress,
				Market: sell.Offer,
			}
			sellerOutput := Output{
				Adress: sell.Adress,
				Main:   buy.Offer,
			}
			return []Output{
				buyerOutput,
				sellerOutput,
			}
		}

	}
	return nil
}
