package main

import "sync_tree/__logs"

func main() {
	__logs.Info("some info piece")
	__logs.Error("asda")
	__logs.Warning("lilt warn")
	__logs.Critical("askjdnhk %v ")
}