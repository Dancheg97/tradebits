package market

import (
	"math/rand"
	"reflect"
	"sync_tree/calc"
	"sync_tree/user"
	"testing"
	"time"
)

func genRandNumber() int {
	rand.Seed(time.Now().UnixNano())
	min := 500
	max := 1000
	randNum := rand.Intn(max-min) + min
	return randNum
}

func TestOperate(t *testing.T) {
	for i := 0; i < 100000; i++ {
		firstRandTrade := Trade{Adress: []byte("a")}
		secondRandTrade := Trade{Adress: []byte("b")}
		randNumbers := []uint64{}
		for i := 0; i < 4; i++ {
			rand.Seed(time.Now().UnixNano())
			min := 1
			max := 100000
			randNum := rand.Intn(max-min) + min
			randNumbers = append(randNumbers, uint64(randNum))
		}
		firstRandTrade.Offer = randNumbers[0]
		firstRandTrade.Recieve = randNumbers[1]
		secondRandTrade.Offer = randNumbers[2]
		secondRandTrade.Recieve = randNumbers[3]
		randBool := rand.Intn(2) != 0
		firstRandTrade.IsSell = randBool
		secondRandTrade.IsSell = !randBool
		trades, outputs := firstRandTrade.operate(secondRandTrade)
		if outputs == nil {
			if !reflect.DeepEqual(firstRandTrade, trades[0]) {
				t.Error("if trades dont operate they should be same")
			}
			if !reflect.DeepEqual(secondRandTrade, trades[1]) {
				t.Error("if trades dont operate they should be the smae")
			}
		} else {
			mainOutputSum := uint64(0)
			marketOutputSum := uint64(0)
			for _, output := range outputs {
				mainOutputSum = mainOutputSum + output.MainOut
				marketOutputSum = marketOutputSum + output.MarketOut
			}
			for _, trade := range trades {
				if trade.IsSell {
					marketOutputSum = marketOutputSum + trade.Offer
				} else {
					mainOutputSum = mainOutputSum + trade.Offer
				}
			}
			if firstRandTrade.IsSell {
				if firstRandTrade.Offer != marketOutputSum {
					t.Error("market sum dont match")
				}
				if secondRandTrade.Offer != mainOutputSum {
					t.Error("main sum dont match")
				}
			} else {
				if firstRandTrade.Offer != mainOutputSum {
					t.Error("main sum dont match")
				}
				if secondRandTrade.Offer != marketOutputSum {
					t.Error("market sum dont match")
				}
			}
		}
	}
}

func TestMarketAddTrade(t *testing.T) {
	Create(adress, name, mesKey, descr, img)
	market := Get(adress)
	for i := 0; i < 15; i++ {
		randNumbers := []uint64{}
		for i := 0; i < 2; i++ {
			randNum := genRandNumber()
			randNumbers = append(randNumbers, uint64(randNum))
		}
		randTrade := Trade{
			Offer:   randNumbers[0],
			Recieve: randNumbers[1],
			IsSell:  rand.Intn(2) != 0,
			Adress:  []byte{byte(rand.Intn(254)), byte(rand.Intn(254))},
		}
		market.OperateTrade(randTrade)
	}
}

/*
This test is creating a sequence of market trade requests, each of which are
then operating on a market, then after each new market request is operated,
this test is checking 'check sum', so that no trade is causing unstable
behaviour. call as main
*/
func TestMultipleTradesOperatingWithCheckSum(t *testing.T) {
	Create(adress, name, mesKey, descr, img)
	mkt := Get(adress)
	checkSumMain := 0
	checkSumMarket := 0
	randomUserAdresses := [][]byte{}
	for i := 0; i < 10000; i++ {
		randAdress := calc.Rand()
		randomUserAdresses = append(randomUserAdresses, randAdress)
		user.Create(randAdress, calc.Rand(), "randtester")
		randTrade := Trade{
			IsSell:  rand.Intn(2) != 0,
			Offer:   uint64(genRandNumber()),
			Recieve: uint64(genRandNumber()),
		}
		if randTrade.IsSell {
			checkSumMarket = checkSumMarket + int(randTrade.Offer)
		} else {
			checkSumMain = checkSumMain + int(randTrade.Offer)
		}
		mkt.OperateTrade(randTrade)
		localCheckSumMain := 0
		localCheckSumMarket := 0
		for _, adr := range randomUserAdresses {
			u := user.Look(adr)
			localCheckSumMain = localCheckSumMain + int(u.Balance)
			localCheckSumMarket = localCheckSumMarket + int(u.MarketBalance(adress))
		}
		for _, trade := range mkt.Buys {
			localCheckSumMain = localCheckSumMain + int(trade.Offer)
		}
		for _, trade := range mkt.Sells {
			localCheckSumMarket = localCheckSumMarket + int(trade.Offer)
		}
		firstCondition := localCheckSumMain == checkSumMain
		secondCondition := localCheckSumMarket == checkSumMarket
		if !(firstCondition && secondCondition) {
			t.Error("error occured")
		}
	}
}
