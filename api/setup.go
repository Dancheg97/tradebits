package api

import (
	"tradebits/api/info"
	"tradebits/api/user"
	"tradebits/crypter"
	"tradebits/mongoer"
	"tradebits/rediser"
)

var marketinfo []byte
var mongo mongoer.IMongoer
var crypt crypter.ICrypter
var redis rediser.IRediser

func Setup(
	marketInfo []byte,
	mongoer mongoer.IMongoer,
	crypter crypter.ICrypter,
	rediser rediser.IRediser,
) {
	user.Setup(marketInfo, mongo, crypt, redis)
	info.Setup(marketInfo, mongo, crypt, redis)
	marketinfo = marketInfo
	mongo = mongoer
	crypt = crypter
	redis = rediser
}
