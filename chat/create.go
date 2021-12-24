package chat

import (
	"bytes"
	"encoding/gob"
	"errors"
	"orb/data"
)

type Message struct {
	message string
}

type chat struct {
	Messages        []Message
	checkedByUser   bool
	checkedByMarket bool
}

func createChat(marketAdress []byte, userAdress []byte) error {
	mktExists := data.Check(marketAdress)
	if !mktExists {
		return errors.New("market does not exist")
	}
	ustExists := data.Check(userAdress)
	if !ustExists {
		return errors.New("user does not exis")
	}
	chatAdress := append(marketAdress, userAdress...)
	chtExists := data.Check(chatAdress)
	if chtExists {
		return errors.New("chat already exists")
	}
	cht := chat{
		Messages: []Message{},
	}
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(cht)
	data.Put(chatAdress, cache.Bytes())
	return nil
}
