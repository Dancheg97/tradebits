package market

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"sync_tree/data"
	"sync_tree/search"
	"sync_tree/trade"

	"github.com/AccelByte/profanity-filter-go"
)


// Create new market by passed values. Checks wether market with passed adress
// exists and creates new one. Here is field description:
//
// - Adress: represens hash of markets public key
//
// - Name: market name visible for users (min 10, max 30)
//
// - MesKey: message key that is gonna be used to check
//
// - Descr: market description visible for users (min 160, max 480)
//
// - Img: url link to the image, bcs market dont store images
//
// - InputFee: fee value, each round number representing 0.01% (min 0, max 500)
//
// - OutputFee: fee value, each round number representing 0.01% (min 0, max 500)
//
// - WorkTime: representing when market is working with messages (min15, max 45)
//
// - Delimiter: value that is representing decimal places of its value (max 10)
func Create(
	adress []byte,
	name string,
	mesKey []byte,
	descr string,
	imgLink string,
	inputFee uint64,
	outputFee uint64,
	workTime string,
	delimiter uint64,
) error {
	if len(adress) != 64 {
		return errors.New("bad adress length")
	}
	if len(name) < 10 || len(name) > 30 {
		return errors.New("bad name length")
	}
	if len(mesKey) < 240 || len(mesKey) > 320 {
		fmt.Println(len(mesKey))
		return errors.New("invalid message key length")
	}
	if len(descr) < 160 || len(descr) > 760 {
		return errors.New("bad description length")
	}
	if inputFee > 500 || outputFee > 500 {
		return errors.New("fee too big")
	}
	if len(workTime) < 10 || len(workTime) > 40 {
		return errors.New("work time is bad")
	}
	if delimiter > 10 {
		return errors.New("delimiter length is too long")
	}
	if data.Check(adress) {
		return errors.New("possibly market already exists")
	}
	if data.Check([]byte(name)) {
		return errors.New("market with that name exists")
	}
	if name[0] == " "[0] || name[len(name)-1] == " "[0] {
		return errors.New("market name start/ends with space")
	}
	isBadName, _, _ := profanityfilter.Filter.ProfanityCheck(name)
	if isBadName {
		return errors.New("name contains bad words")
	}
	isBadDescr, words, _ := profanityfilter.Filter.ProfanityCheck(descr)
	if isBadDescr {
		errStr := fmt.Sprint("description contains profane words", words)
		return errors.New(errStr)
	}
	data.Put([]byte(name), []byte{})
	pool := trade.TradePool{
		Buys:    []trade.Buy{},
		Sells:   []trade.Sell{},
		Outputs: []trade.Output{},
	}
	newMarket := market{
		adress:    adress,
		Name:      name,
		Descr:     descr,
		Img:       imgLink,
		MesKey:    mesKey,
		OpCount:   0,
		Pool:      pool,
		InputFee:  inputFee,
		OutputFee: outputFee,
		WorkTime:  workTime,
		Delimiter: delimiter,
		Users:     [][]byte{},
	}
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(newMarket)
	data.Put(adress, cache.Bytes())
	search.Add(name, adress)
	return nil
}
