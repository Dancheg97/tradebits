package market

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var adress = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 63}
var mesKey = []byte{1, 2, 3, 4, 5}
var img = "asset image link . example"
var name = "newAsset"
var descr = "descrx"

func TestCreateNewMarket(t *testing.T) {
	err := Create(adress, name, mesKey, descr, img)
	if err != nil {
		t.Error("failed to craete new market")
		return
	}
}

func TestCreateExistingMarket(t *testing.T) {
	err := Create(adress, name, mesKey, descr, img)
	if err != nil {
		return
	}
	t.Error("failed to craete new market")
}

func TestGetFreeAssetFromDB(t *testing.T) {
	market := Get(adress)
	defer market.Save()
	if reflect.DeepEqual(market.MesKey, mesKey) {
		return
	}
	t.Error("keys are not the same, get asset error")
}

func TestGetBusyAssetFromDB(t *testing.T) {
	freeAsset := Get(adress)
	defer freeAsset.Save()
	busyAsset := Get(adress)
	if busyAsset != nil {
		t.Error("attempt to get busy asset from db")
		return
	}
}

func TestAssetLook(t *testing.T) {
	market := Look(adress)
	if reflect.DeepEqual(market.MesKey, mesKey) {
		return
	}
	t.Error("keys are not the same, look asset error")
}

func TestMarketAddTrade(t *testing.T) {
	Create(adress, name, mesKey, descr, img)
	market := Get(adress)
	for i := 0; i < 15; i++ {
		randNumbers := []uint64{}
		for i := 0; i < 2; i++ {
			rand.Seed(time.Now().UnixNano())
			min := 1
			max := 100
			randNum := rand.Intn(max-min) + min
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

func TestMultipleTradesOperatingWithCheckSum(t *testing.T) {
	
}

