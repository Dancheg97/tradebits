package user

import (
	"tradebits/crypter"
	"tradebits/mongoer"
	"tradebits/rediser"
)

var mongo mongoer.IMongoer
var crypt crypter.ICrypter
var redis rediser.IRediser

func Setup(
	mongoer mongoer.IMongoer,
	crypter crypter.ICrypter,
	rediser rediser.IRediser,
) {
	mongo = mongoer
	crypt = crypter
	redis = rediser
}
