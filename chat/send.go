package chat

import "orb/user"

func UserSend(message string, marketAdress []byte, userAdress []byte) {
	user := user.Get(userAdress)
	
}

func MarketSend(message string, marketAdress []byte, userAdress []byte) {
	
}