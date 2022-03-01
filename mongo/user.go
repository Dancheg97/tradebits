package mongo

type User struct {
	PubKey   string   `json:"PubKey" bson:"PubKey"`
	Balance  int      `json:"Balance" bson:"Balance"`
	Messages []string `json:"Messages" bson:"Messages"`
}

// checks wether user exists in mongo
func UserCheck(pubkey string) (bool, error) {
	return true, nil
}

// creates new user in mongo
func UserCreate(pubkey string) error {
	return nil
}

// put new message to related to user
func UserPutMessage(pubkey string, message string) error {
	return nil
}

// increases user balance by some amount
func UserIncreaseBalance(pubkey string, increase int) error {
	return nil
}

// decrease user balance by some amount
func UserDecreateBalance(pubkey string, decrease int) error {
	return nil
}
