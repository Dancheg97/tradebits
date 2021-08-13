package trade

import (
	"reflect"
	"testing"
)

func TestMatchCloseBuy(t *testing.T) {
	buy := Buy{
		Adress:  []byte{0},
		Offer:   60,
		Recieve: 40,
	}
	sell := Sell{
		Adress:  []byte{1},
		Offer:   100,
		Recieve: 100,
	}

	outputs := buy.match(&sell)

	expectedBuyerOutput := Output{
		Adress: []byte{0},
		Market: 40,
	}
	expectedSellerOutput := Output{
		Adress: []byte{1},
		Main:   60,
	}
	expectedBuy := Buy{
		Adress:  []byte{0},
		Offer:   0,
		Recieve: 0,
	}
	expectedSell := Sell{
		Adress:  []byte{1},
		Offer:   60,
		Recieve: 40,
	}
	if !reflect.DeepEqual(outputs[0], expectedBuyerOutput) {
		t.Error("buyer output not matching", outputs[0], expectedBuyerOutput)
	}
	if !reflect.DeepEqual(outputs[1], expectedSellerOutput) {
		t.Error("seller output not matching", outputs[1], expectedSellerOutput)
	}
	if !reflect.DeepEqual(buy, expectedBuy) {
		t.Error("buy is not matching", buy, expectedBuy)
	}
	if !reflect.DeepEqual(sell, expectedSell) {
		t.Error("sell is not matching", sell, expectedSell)
	}
}

func TestMatchCloseSell(t *testing.T) {
	buy := Buy{
		Adress:  []byte{0},
		Offer:   100,
		Recieve: 100,
	}
	sell := Sell{
		Adress:  []byte{1},
		Offer:   60,
		Recieve: 40,
	}
	outputs := buy.match(&sell)

	expectedBuyerOutput := Output{
		Adress: []byte{0},
		Market: 60,
	}
	expectedSellerOutput := Output{
		Adress: []byte{1},
		Main:   40,
	}
	expectedBuy := Buy{
		Adress:  []byte{0},
		Offer:   60,
		Recieve: 40,
	}
	expectedSell := Sell{
		Adress:  []byte{1},
		Offer:   0,
		Recieve: 0,
	}
	if !reflect.DeepEqual(outputs[0], expectedBuyerOutput) {
		t.Error("buyer output not matching", outputs[0], expectedBuyerOutput)
	}
	if !reflect.DeepEqual(outputs[1], expectedSellerOutput) {
		t.Error("seller output not matching", outputs[1], expectedSellerOutput)
	}
	if !reflect.DeepEqual(buy, expectedBuy) {
		t.Error("buy is not matching", buy, expectedBuy)
	}
	if !reflect.DeepEqual(sell, expectedSell) {
		t.Error("sell is not matching", sell, expectedSell)
	}
}

func TestMatchBothOfferTooMuch(t *testing.T) {
	buy := Buy{
		Adress:  []byte{0},
		Offer:   120,
		Recieve: 100,
	}
	sell := Sell{
		Adress:  []byte{1},
		Offer:   110,
		Recieve: 90,
	}

	outputs := buy.match(&sell)

	expectedBuyerOutput := Output{
		Adress: []byte{0},
		Market: 110,
	}
	expectedSellerOutput := Output{
		Adress: []byte{1},
		Main:   120,
	}
	expectedBuy := Buy{
		Adress:  []byte{0},
		Offer:   0,
		Recieve: 0,
	}
	expectedSell := Sell{
		Adress:  []byte{1},
		Offer:   0,
		Recieve: 0,
	}
	if !reflect.DeepEqual(outputs[0], expectedBuyerOutput) {
		t.Error("buyer output not matching", outputs[0], expectedBuyerOutput)
	}
	if !reflect.DeepEqual(outputs[1], expectedSellerOutput) {
		t.Error("seller output not matching", outputs[1], expectedSellerOutput)
	}
	if !reflect.DeepEqual(buy, expectedBuy) {
		t.Error("buy is not matching", buy, expectedBuy)
	}
	if !reflect.DeepEqual(sell, expectedSell) {
		t.Error("sell is not matching", sell, expectedSell)
	}
}
func TestMatchBuyOffersMuchOnSameAmountOfSellRecieve(t *testing.T) {
	buy := Buy{
		Adress:  []byte{0},
		Offer:   200,
		Recieve: 100,
	}
	sell := Sell{
		Adress:  []byte{1},
		Offer:   100,
		Recieve: 100,
	}
	outputs := buy.match(&sell)
	expectedBuyerOutput := Output{
		Adress: []byte{0},
		Market: 100,
	}
	expectedSellerOutput := Output{
		Adress: []byte{1},
		Main:   200,
	}
	expectedBuy := Buy{
		Adress:  []byte{0},
		Offer:   0,
		Recieve: 0,
	}
	expectedSell := Sell{
		Adress:  []byte{1},
		Offer:   0,
		Recieve: 0,
	}
	if !reflect.DeepEqual(outputs[0], expectedBuyerOutput) {
		t.Error("buyer output not matching", outputs[0], expectedBuyerOutput)
	}
	if !reflect.DeepEqual(outputs[1], expectedSellerOutput) {
		t.Error("seller output not matching", outputs[1], expectedSellerOutput)
	}
	if !reflect.DeepEqual(buy, expectedBuy) {
		t.Error("buy is not matching", buy, expectedBuy)
	}
	if !reflect.DeepEqual(sell, expectedSell) {
		t.Error("sell is not matching", sell, expectedSell)
	}
}

func TestMatchSellOffersMuchOnSameAmountOfBuyRecieve(t *testing.T) {
	buy := Buy{
		Adress:  []byte{0},
		Offer:   100,
		Recieve: 100,
	}
	sell := Sell{
		Adress:  []byte{1},
		Offer:   200,
		Recieve: 100,
	}
	outputs := buy.match(&sell)

	expectedBuyerOutput := Output{
		Adress: []byte{0},
		Market: 200,
	}
	expectedSellerOutput := Output{
		Adress: []byte{1},
		Main:   100,
	}
	expectedBuy := Buy{
		Adress:  []byte{0},
		Offer:   0,
		Recieve: 0,
	}
	expectedSell := Sell{
		Adress:  []byte{1},
		Offer:   0,
		Recieve: 0,
	}
	if !reflect.DeepEqual(outputs[0], expectedBuyerOutput) {
		t.Error("buyer output not matching", outputs[0], expectedBuyerOutput)
	}
	if !reflect.DeepEqual(outputs[1], expectedSellerOutput) {
		t.Error("seller output not matching", outputs[1], expectedSellerOutput)
	}
	if !reflect.DeepEqual(buy, expectedBuy) {
		t.Error("buy is not matching", buy, expectedBuy)
	}
	if !reflect.DeepEqual(sell, expectedSell) {
		t.Error("sell is not matching", sell, expectedSell)
	}
}

func TestMatchCloseBoth(t *testing.T) {
	buy := Buy{
		Adress:  []byte{0},
		Offer:   100,
		Recieve: 50,
	}
	sell := Sell{
		Adress:  []byte{1},
		Offer:   50,
		Recieve: 100,
	}

	outputs := buy.match(&sell)

	expectedBuyerOutput := Output{
		Adress: []byte{0},
		Market: 50,
	}
	expectedSellerOutput := Output{
		Adress: []byte{1},
		Main:   100,
	}
	expectedBuy := Buy{
		Adress:  []byte{0},
		Offer:   0,
		Recieve: 0,
	}
	expectedSell := Sell{
		Adress:  []byte{1},
		Offer:   0,
		Recieve: 0,
	}
	if !reflect.DeepEqual(outputs[0], expectedBuyerOutput) {
		t.Error("buyer output not matching", outputs[0], expectedBuyerOutput)
	}
	if !reflect.DeepEqual(outputs[1], expectedSellerOutput) {
		t.Error("seller output not matching", outputs[1], expectedSellerOutput)
	}
	if !reflect.DeepEqual(buy, expectedBuy) {
		t.Error("buy is not matching", buy, expectedBuy)
	}
	if !reflect.DeepEqual(sell, expectedSell) {
		t.Error("sell is not matching", sell, expectedSell)
	}
}

func TestNotMatch(t *testing.T) {
	buy := Buy{
		Adress:  []byte{0},
		Offer:   50,
		Recieve: 100,
	}
	sell := Sell{
		Adress:  []byte{1},
		Offer:   50,
		Recieve: 100,
	}

	outputs := buy.match(&sell)

	expectedBuy := Buy{
		Adress:  []byte{0},
		Offer:   50,
		Recieve: 100,
	}
	expectedSell := Sell{
		Adress:  []byte{1},
		Offer:   50,
		Recieve: 100,
	}
	if len(outputs) != 0 {
		t.Error("there should not be any outputs in non matching trades")
	}
	if !reflect.DeepEqual(buy, expectedBuy) {
		t.Error("buy is not matching", buy, expectedBuy)
	}
	if !reflect.DeepEqual(sell, expectedSell) {
		t.Error("sell is not matching", sell, expectedSell)
	}
}

func TestAddSingleBuyToPool(t *testing.T) {
	tp := TradePool{
		Buys:    []Buy{},
		Sells:   []Sell{},
		Outputs: []Output{},
	}
	buy := Buy{
		Adress:  []byte{0},
		Offer:   50,
		Recieve: 100,
	}

	tp.OperateBuy(buy)

	if len(tp.Buys) != 1 {
		t.Error("There should be one active buy order on the market")
	}
	if len(tp.Sells) != 0 {
		t.Error("There should not be any active sells on the market")
	}
	if len(tp.Outputs) != 0 {
		t.Error("There should not be any active outputs on the market")
	}
	if !reflect.DeepEqual(buy, tp.Buys[0]) {
		t.Error("Trade in pool is not equal to the added")
	}
}

func TestSortingOnAddingMultipleBuys(t *testing.T) {
	tp := TradePool{
		Buys:    []Buy{},
		Sells:   []Sell{},
		Outputs: []Output{},
	}
	bestBuy := Buy{
		Adress:  []byte{0},
		Offer:   100,
		Recieve: 50,
	}
	midBuy := Buy{
		Adress:  []byte{0},
		Offer:   75,
		Recieve: 75,
	}
	worstBuy := Buy{
		Adress:  []byte{0},
		Offer:   50,
		Recieve: 100,
	}
	tp.OperateBuy(midBuy)
	tp.OperateBuy(bestBuy)
	tp.OperateBuy(worstBuy)

	if len(tp.Buys) != 3 {
		t.Error("there should be 3 active buys in trade pool")
	}
	if len(tp.Sells) != 0 {
		t.Error("there should not be any active sells in trade pool")
	}
	if len(tp.Outputs) != 0 {
		t.Error("outputs should not appear in pool after adding only buys")
	}
	if !reflect.DeepEqual(tp.Buys[0], bestBuy) {
		t.Error("the first buy in pool should be the best one")
	}
	if !reflect.DeepEqual(tp.Buys[1], midBuy) {
		t.Error("the average trade should be mid one")
	}
	if !reflect.DeepEqual(tp.Buys[2], worstBuy) {
		t.Error("the last trade should be the worst one")
	}
}

func TestAddSingleSellToPool(t *testing.T) {
	tp := TradePool{
		Sells:   []Sell{},
		Buys:    []Buy{},
		Outputs: []Output{},
	}
	sell := Sell{
		Adress:  []byte{1},
		Offer:   50,
		Recieve: 100,
	}

	tp.OperateSell(sell)

	if len(tp.Buys) != 0 {
		t.Error("there should not be any active buys in pool")
	}
	if len(tp.Sells) != 1 {
		t.Error("there should be a single sell order in pool")
	}
	if len(tp.Outputs) != 0 {
		t.Error("there should not be any outputs in current")
	}
	if !reflect.DeepEqual(tp.Sells[0], sell) {
		t.Error("the current order is not matching with added one")
	}
}

func TestAddMultipleSells(t *testing.T) {
	tp := TradePool{
		Sells:   []Sell{},
		Buys:    []Buy{},
		Outputs: []Output{},
	}
	bestSell := Sell{
		Adress:  []byte{0},
		Offer:   100,
		Recieve: 50,
	}
	midSell := Sell{
		Adress:  []byte{0},
		Offer:   75,
		Recieve: 75,
	}
	worstSell := Sell{
		Adress:  []byte{0},
		Offer:   50,
		Recieve: 100,
	}

	tp.OperateSell(midSell)
	tp.OperateSell(bestSell)
	tp.OperateSell(worstSell)

	if len(tp.Sells) != 3 {
		t.Error("there should be 3 active sells in trade pool")
	}
	if len(tp.Buys) != 0 {
		t.Error("there should not be any active buys in trade pool")
	}
	if len(tp.Outputs) != 0 {
		t.Error("outputs should not appear in pool after adding only buys")
	}
	if !reflect.DeepEqual(tp.Sells[0], bestSell) {
		t.Error("the first sell in pool should be the best one")
	}
	if !reflect.DeepEqual(tp.Sells[1], midSell) {
		t.Error("the average trade should be mid one")
	}
	if !reflect.DeepEqual(tp.Sells[2], worstSell) {
		t.Error("the last trade in pool should be the worst one")
	}
}

func checkIfElementIsMissing(slice []Output, elem Output) bool {
	for _, v := range slice {
		if reflect.DeepEqual(elem, v) {
			return false
		}
	}
	return true
}

func TestOperateMultipleBuysAndSells(t *testing.T) {
	tp := TradePool{
		Sells:   []Sell{},
		Buys:    []Buy{},
		Outputs: []Output{},
	}
	bestSell := Sell{
		Adress:  []byte{0},
		Offer:   100,
		Recieve: 50,
	}
	midSell := Sell{
		Adress:  []byte{1},
		Offer:   75,
		Recieve: 75,
	}
	worstSell := Sell{
		Adress:  []byte{2},
		Offer:   50,
		Recieve: 100,
	}
	bestBuy := Buy{
		Adress:  []byte{3},
		Offer:   100,
		Recieve: 50,
	}
	midBuy := Buy{
		Adress:  []byte{4},
		Offer:   75,
		Recieve: 75,
	}
	worstBuy := Buy{
		Adress:  []byte{5},
		Offer:   50,
		Recieve: 100,
	}

	tp.OperateBuy(bestBuy)
	tp.OperateSell(worstSell)
	tp.OperateBuy(midBuy)
	tp.OperateSell(midSell)
	tp.OperateSell(bestSell)
	tp.OperateBuy(worstBuy)

	firstOutput := Output{
		Adress: []byte{0},
		Main:   50,
	}
	secondOutput := Output{
		Adress: []byte{1},
		Main:   75,
	}
	thirdOutput := Output{
		Adress: []byte{2},
		Main:   100,
	}
	fourthOutput := Output{
		Adress: []byte{3},
		Market: 50,
	}
	fifthOutput := Output{
		Adress: []byte{4},
		Market: 75,
	}
	sixthOutput := Output{
		Adress: []byte{5},
		Market: 100,
	}

	allOutputs := []Output{
		firstOutput,
		secondOutput,
		thirdOutput,
		fourthOutput,
		fifthOutput,
		sixthOutput,
	}

	if len(tp.Outputs) != 6 {
		t.Error("output should exist for every user")
	}
	for idx, output := range allOutputs {
		if checkIfElementIsMissing(tp.Outputs, output) {
			t.Error("element by index ", idx, " is missing")
		}
	}
}

func TestAddingNonMatchingTrades(t *testing.T) {
	tp := TradePool{
		Sells:   []Sell{},
		Buys:    []Buy{},
		Outputs: []Output{},
	}
	firstSell := Sell{
		Adress:  []byte{0},
		Offer:   50,
		Recieve: 100,
	}
	secondSell := Sell{
		Adress:  []byte{0},
		Offer:   75,
		Recieve: 100,
	}
	firstBuy := Buy{
		Adress:  []byte{3},
		Offer:   50,
		Recieve: 100,
	}
	secondBuy := Buy{
		Adress:  []byte{3},
		Offer:   75,
		Recieve: 100,
	}

	tp.OperateBuy(firstBuy)
	tp.OperateSell(firstSell)
	tp.OperateBuy(secondBuy)
	tp.OperateSell(secondSell)

	if len(tp.Sells) != 2 {
		t.Error("there should be 2 active sells on current market")
	}
	if len(tp.Buys) != 2 {
		t.Error("there should be 2 active buys on current market")
	}
	if !reflect.DeepEqual(tp.Buys[0], secondBuy) {
		t.Error("first pool buy is better, should be equal to secondBuy")
	}
	if !reflect.DeepEqual(tp.Buys[1], firstBuy) {
		t.Error("second pool buy is worse, should be equal to firstBuy")
	}
	if !reflect.DeepEqual(tp.Sells[0], secondSell) {
		t.Error("first pool buy is better, should be equal to secondBuy")
	}
	if !reflect.DeepEqual(tp.Sells[1], firstSell) {
		t.Error("second pool sell is worse, should be  equal to firstSell")
	}
}

func TestSellClosedMultipleBuysInPool(t *testing.T) {
	tp := TradePool{
		Sells:   []Sell{},
		Buys:    []Buy{},
		Outputs: []Output{},
	}
	bigSell := Sell{
		Adress:  []byte{0},
		Offer:   200,
		Recieve: 200,
	}
	firstSmallBuy := Buy{
		Adress:  []byte{1},
		Offer:   90,
		Recieve: 90,
	}
	secondSmallBuy := Buy{
		Adress:  []byte{2},
		Offer:   110,
		Recieve: 110,
	}

	tp.OperateBuy(firstSmallBuy)
	tp.OperateBuy(secondSmallBuy)
	tp.OperateSell(bigSell)

	firstOutput := Output{
		Adress: []byte{0},
		Main:   90,
	}
	secondOutput := Output{
		Adress: []byte{0},
		Main:   110,
	}
	thirdOutput := Output{
		Adress: []byte{1},
		Market: 90,
	}
	fourthOutput := Output{
		Adress: []byte{2},
		Market: 110,
	}
	if len(tp.Outputs) != 4 {
		t.Error("there should be 4 outputs in pool")
	}
	if checkIfElementIsMissing(tp.Outputs, firstOutput) {
		t.Error("cant find first expected output")
	}
	if checkIfElementIsMissing(tp.Outputs, secondOutput) {
		t.Error("cant find second expected output")
	}
	if checkIfElementIsMissing(tp.Outputs, thirdOutput) {
		t.Error("cant find third expected output")
	}
	if checkIfElementIsMissing(tp.Outputs, fourthOutput) {
		t.Error("cant find fourth expected output")
	}
}
