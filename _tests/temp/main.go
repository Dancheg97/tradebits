package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3}
	s = s[1:]
	fmt.Println(s)
}
