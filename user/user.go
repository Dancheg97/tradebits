package user

import (
	"bc_server/lock"
	"encoding/json"
	"errors"
	"runtime"
)

type User struct {
	MainBalance   uint64            `json:"MainBalance"`
	MessageKey    []byte            `json:"MessageKey"`
	Image         []byte            `json:"Image"`
	AssetBalances map[string]uint64 `json:"AssetBalances"`
	adress        []byte
}

func NewUser(adress []byte) (User, error) {
	lock.Lock(adress)
	user := User{}
	user.adress = adress
	_, dbErr := userDB.Get(adress, nil)
	if dbErr == nil {
		return user, errors.New("user already exsits")
	}
	user.MainBalance = 0
	user.MessageKey = nil
	user.Image = nil
	user.AssetBalances = make(map[string]uint64)
	userAsBytes, _ := json.Marshal(&user)
	userDB.Put(adress, userAsBytes, nil)
	runtime.SetFinalizer(
		&user,
		user.unlcok,
	)
	return user, nil
}

func GetUser(adress []byte) (User, error) {
	user := User{}
	user.adress = adress
	userBytes, dbErr := userDB.Get(adress, nil)
	if dbErr != nil {
		return user, errors.New("user do not exist")
	}
	json.Unmarshal(userBytes, &user)
	return user, nil
}

func (user User) unlcok() {
	lock.Unlock(user.adress)
}

func (user User) SetMainBalance(balance uint64) {
	user.MainBalance = balance
	userAsBytes, _ := json.Marshal(&user)
	userDB.Put(user.adress, userAsBytes, nil)
}

func (user User) SetMessageKey(messageKey []byte) {
	user.MessageKey = messageKey
	userAsBytes, _ := json.Marshal(&user)
	userDB.Put(user.adress, userAsBytes, nil)
}

func (user User) SetImage(image []byte) {
	user.Image = image
	userAsBytes, _ := json.Marshal(&user)
	userDB.Put(user.adress, userAsBytes, nil)
}

func (user User) SetAssetBalance(asset []byte, balance uint64) {
	user.AssetBalances[string(asset)] = balance
	userAsBytes, _ := json.Marshal(&user)
	userDB.Put(user.adress, userAsBytes, nil)
}
