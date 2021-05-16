package database

import (
	"encoding/json"
	"errors"
	"github.com/syndtr/goleveldb/leveldb"
)

type User struct {
	MainBalance       uint64            `json:"MainBalance"`
	MessageKey        []byte            `json:"MessageKey"`
	Image             []byte            `json:"Image"`
	ExchangerBalances map[string]uint64 `json:"ExchangerBalances"`
	adress            []byte
}

var UserDB, _ = leveldb.OpenFile("database/userData", nil)

func NewUser(adress []byte) (User, error) {
	user := User{}
	user.adress = adress
	_, dbErr := UserDB.Get(adress, nil)
	if dbErr == nil {
		return user, errors.New("user already exsits")
	}
	user.MainBalance = 0
	user.MessageKey = nil
	user.Image = nil
	user.ExchangerBalances = make(map[string]uint64)
	userAsBytes, _ := json.Marshal(&user)
	UserDB.Put(adress, userAsBytes, nil)
	return user, nil
}

func GetUser(adress []byte) (User, error) {
	user := User{}
	user.adress = adress
	userBytes, dbErr := UserDB.Get(adress, nil)
	if dbErr != nil {
		return user, errors.New("user do not exist")
	}
	json.Unmarshal(userBytes, &user)
	return user, nil
}

func (user User) SetMainBalance(balance uint64) {
	user.MainBalance = balance
	userAsBytes, _ := json.Marshal(&user)
	UserDB.Put(user.adress, userAsBytes, nil)
}

func (user User) SetMessageKey(messageKey []byte) {
	user.MessageKey = messageKey
	userAsBytes, _ := json.Marshal(&user)
	UserDB.Put(user.adress, userAsBytes, nil)
}

func (user User) SetImage(image []byte) {
	user.Image = image
	userAsBytes, _ := json.Marshal(&user)
	UserDB.Put(user.adress, userAsBytes, nil)
}

func (user User) SetExchangerBalance(exchanger []byte, balance uint64) {
	user.ExchangerBalances[string(exchanger)] = balance
	userAsBytes, _ := json.Marshal(&user)
	UserDB.Put(user.adress, userAsBytes, nil)
}
