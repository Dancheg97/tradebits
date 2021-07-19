package main

import "fmt"

func main() {
	firstArr := []byte{0, 1, 2, 3, 4}
	fmt.Println(firstArr)
	secondArr := append(firstArr[:4], firstArr[4+1:]...)
	fmt.Println(secondArr)
}
