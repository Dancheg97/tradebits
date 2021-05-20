package main

import (
	"bc_server/database"
	"errors"
	"fmt"
)

func testNewUserGeneration() error {
	key := []byte{1, 2, 3}
	user, generror := database.NewUser(key)
	if generror != nil {
		return generror
	}
	user, genTwiceError := database.NewUser(key)
	if genTwiceError == nil {
		return errors.New("error user with same id")
	}
	if user.MainBalance != 0 {
		return errors.New("generated non zero balance")
	}
	if user.Image != nil {
		return errors.New("generated non empty image")
	}
	return nil
}

func getUserTest() error {
	key := []byte{1, 2, 3}
	user, _ := database.NewUser(key)
	user, getErr := database.GetUser(key)
	if getErr != nil {
		return errors.New("get user err")
	}
	if user.MainBalance != 0 {
		return errors.New("user data err")
	}
	return nil
}

func setUserParamsTest() error {
	key := []byte{1, 2, 3}
	user, _ := database.NewUser(key)
	user.SetMainBalance(12)
	sameUser, _ := database.GetUser(key)
	if sameUser.MainBalance != 12 {
		return errors.New("bad balance")
	}
	return nil
}

func main() {
	tests := []error{
		testNewUserGeneration(),
		getUserTest(),
		setUserParamsTest(),
	}
	for index, err := range tests {
		fmt.Println("test", index, "error:", err)
	}
}
