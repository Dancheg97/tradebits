package database

import (
	"errors"
)

type User struct {
	imageUUID            []byte
	messageKeyUUID       []byte
	mainBalanceUUID      []byte
	exchangerBalanceUUID []byte
}

// func (user User) newRandom() {
// 	gou
// }

func (user User) toBytes() []byte {
	return append(
		append(user.imageUUID, user.mainBalanceUUID...),
		append(user.mainBalanceUUID, user.messageKeyUUID...)...,
	)
}

func new() User {
	return User{imageUUID: gouuid.New()}
}

func WriteNewUser(adress []byte) error {
	_, err := db.Get(adress, nil)
	if err != nil {
		return errors.New("user exists")
	}
	db.Put(adress, nil, nil)
	return nil
}

func WriteUserMainBalance() {

}

func WriteUserImage() {

}

func WriteUserMesageKey() {

}

func WriteUserExchangerBalance() {

}
