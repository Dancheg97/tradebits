package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filebytes, _ := ioutil.ReadFile("adr.pem")
	fmt.Println(string(filebytes))
}
