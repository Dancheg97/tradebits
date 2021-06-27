package main

import "fmt"


func main() {
	m := make(map[string]string)
	m["123"] = m["123"] + "some stuff"
	fmt.Println(m["123"])
}
