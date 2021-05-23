package main

import (
	"sync_tree/user"
)

func testCreateUser() {
	adress := []byte{1, 2, 3}
	mesKey := []byte{1, 2, 3, 4, 5}
	img := "user image link"
	user.Create(adress, mesKey, img)
}

func main() {
	testCreateUser()
}
