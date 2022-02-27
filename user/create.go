package user

import (
	"bytes"
	"encoding/gob"
	"errors"
	"orb/data"

	filter "github.com/AccelByte/profanity-filter-go"
)

type user struct {
	adress   []byte
	Balance  uint64
	MesKey   []byte
	Name     string
	Balances map[string]uint64
	Messages map[string][][]byte
	Arch     map[string]string
}

type 

/*
Create new user, in case there is already user with same adress
the error will be logged
*/
func Create(adress []byte, mesKey []byte, name string) error {
	if len(adress) != 64 {
		return errors.New("bad adress length")
	}
	if len(name) > 12 {
		return errors.New("name too big")
	}
	if len(mesKey) < 240 || len(mesKey) > 320 {
		return errors.New("invalid message key length")
	}
	if data.Check(adress) {
		return errors.New("possibly user already exists")
	}
	if data.Check([]byte(name)) {
		return errors.New("user with that name exists")
	}
	if name[0] == " "[0] || name[len(name)-1] == " "[0] {
		return errors.New("market name start/ends with space")
	}
	isBadName, _, _ := filter.Filter.ProfanityCheck(name)
	if isBadName {
		return errors.New("name contains bad words")
	}
	data.Put([]byte(name), []byte{})
	u := user{
		Balance:  0,
		MesKey:   mesKey,
		Name:     name,
		Balances: make(map[string]uint64),
		Messages: make(map[string][][]byte),
		Arch:     map[string]string{},
	}
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(u)
	data.Put(adress, cache.Bytes())
	return nil
}
