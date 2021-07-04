package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirName := path.Dir(filename)
	fmt.Println(dirName)
}
