package main

import (
	"encoding/json"
	"log"
	"os"
	"tradebits/crypter"
	"tradebits/graylog"
	"tradebits/info"
	"tradebits/market"
	"tradebits/mongoer"
	"tradebits/operator"
	"tradebits/rediser"
	"tradebits/user"

	"github.com/joho/godotenv"
)

var redis_client rediser.IRediser
var mongo_client mongoer.IMongoer
var crypt_client crypter.ICrypter
var mkt_info []byte

func initGraylog(ch chan<- error) {
	ch <- graylog.Setup(readConfigField("GRAYLOG_API"), 60)
}

func initRedis(ch chan<- error) {
	rds, err := rediser.Get(readConfigField("REDIS_HOST"))
	ch <- err
	redis_client = rds
}

func initCrypt(ch chan<- error) {
	crp, err := crypter.Get(readConfigField("MARKET_PRIVATEKEY"))
	ch <- err
	crypt_client = crp
	initInfo(ch)
}

func initMongo(ch chan<- error) {
	mongo, err := mongoer.Get(readConfigField("MONGO_HOST"))
	ch <- err
	ch <- mongo.CreateCollection("user")
	ch <- mongo.CreateCollection("net")
	ch <- mongo.CreateCollection("trades")
	ch <- mongo.CreateIndex("user", "key", "hashed")
	ch <- mongo.CreateIndex("trades", "ukey", "hashed")
	ch <- mongo.CreateIndex("trades", "mkey", "hashed")
}

func initInfo(ch chan<- error) {
	m := map[string]string{
		"name":      readConfigField("MARKET_NAME"),
		"mkey":      crypt_client.Pub(),
		"descr":     readConfigField("MARKET_DESCR"),
		"img":       readConfigField("MARKET_IMG"),
		"worktime":  readConfigField("MARKET_WORKTIME"),
		"fee":       readConfigField("MARKET_FEE"),
		"delimiter": readConfigField("MARKET_DELIMITER"),
	}
	inf, err := json.Marshal(m)
	mkt_info = inf
	ch <- err
}

func init() {
	godotenv.Load()
	setchan := make(chan error)
	go initGraylog(setchan)
	go initRedis(setchan)
	go initCrypt(setchan)
	go initMongo(setchan)
	for i := 0; i < 11; i++ {
		err := <-setchan
		if err != nil {
			log.Fatal(err)
		}
	}
	info.Setup(mkt_info, mongo_client, crypt_client, redis_client)
	market.Setup(mkt_info, mongo_client, crypt_client, redis_client)
	operator.Setup(mkt_info, mongo_client, crypt_client, redis_client)
	user.Setup(mongo_client, crypt_client, redis_client)
	log.Println("Setup sucess...")
}

func readConfigField(field string) string {
	envar, found := os.LookupEnv(field)
	if !found {
		log.Fatal("problem loading .ENV field: ", envar)
	}
	return envar
}
