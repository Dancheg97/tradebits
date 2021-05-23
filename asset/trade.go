package asset

type trade struct {
	adress    []byte
	Direction bool
	Id        []byte
	Maker     []byte
	Offer     uint64
	Course    uint64
}
