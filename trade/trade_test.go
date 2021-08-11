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
