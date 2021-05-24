package asset

type market struct {
	adress []byte
	trades map[string]trade
}

func create()