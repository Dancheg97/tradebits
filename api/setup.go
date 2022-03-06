package api

import (
	"tradebits/crypter"
	"tradebits/mongoer"
	"tradebits/rediser"
)

var info []byte
var mongo mongoer.IMongoer
var crypt crypter.ICrypter
var redis rediser.IRediser

func Setup(
	marketInfo []byte,
	mongoer mongoer.IMongoer,
	crypter crypter.ICrypter,
	rediser rediser.IRediser,
) {
	info = marketInfo
	mongo = mongoer
	crypt = crypter
	redis = rediser
}
