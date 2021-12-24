package chat

import "orb/user"

func UserSend(message string, marketAdress []byte, userAdress []byte) {
	user := user.Get(userAdress)
	
}
