package node

import ()

type node struct {
	privKey     []byte
	pubKey      []byte
	ownerAdress []byte
	connections []string
}
