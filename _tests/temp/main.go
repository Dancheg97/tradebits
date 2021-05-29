package main

import (
	"os"
)

func main() {
	os.RemoveAll("data/data")
	os.RemoveAll("market/data")
	os.RemoveAll("user/data")
}
