package market

import "sync"

/*

*/


type search struct {
	markets map[string][]byte
	mu      sync.Mutex
}

var searchMap = make(map[string][]byte)

func Save() {

}

func Load() {

}

func Find() {

}
