package main

import (
	"fmt"
)

func main() {
	firstArr := []byte{0, 1, 2, 3, 4}
	for i, _ := range firstArr {
		fmt.Println(i)
	}
}
