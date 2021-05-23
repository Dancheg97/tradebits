package asset

type trade struct {
	Direction bool
	Id        []byte
	Maker     []byte
	Offer     uint64
	Course    uint64
}

