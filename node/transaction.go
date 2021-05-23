package node

import (
	"bytes"
	"encoding/gob"
	"sync_tree/__logs"
	"sync_tree/_data"
	"sync_tree/_lock"
)

type node struct {
	adress 
}
