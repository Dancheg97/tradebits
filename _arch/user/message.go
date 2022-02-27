package user

/*
Function to add message from some adress to concrete user
*/
func (u *user) PutMessage(adress []byte, mes []byte) {
	strAdress := string(adress)
	if u.Messages[strAdress] == nil {
		u.Messages[strAdress] = [][]byte{mes}
		return
	}
	u.Messages[strAdress] = append(u.Messages[strAdress], mes)
}

/*
This function is made to get all new messages and to put all current messages
to archieve
*/
func (u *user) GetMessages(adress []byte) [][]byte {
	return u.Messages[string(adress)]
}
