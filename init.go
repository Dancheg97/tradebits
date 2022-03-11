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

func initGraylog() {
	err := graylog.Setup(readConfigField("GRAYLOG_API"), 60)
	if err != nil {
		log.Panic("error connecting to graylog")
	}
	log.Println("Connected to graylog")
}

func initRedis(ch chan<- rediser.IRediser) {
	rds, err := rediser.Get(readConfigField("REDIS_HOST"))
	if err != nil {
		log.Panic("Unable to connect to Redis")
	}
	log.Println("Connected to redis")
	ch <- rds
}

func initCrypt(ch chan<- crypter.ICrypter) {
	crypter, err := crypter.Get(readConfigField("MARKET_PRIVATEKEY"))
	if err != nil {
		log.Panic("Unable to prepare cryptography for crypt module")
	}
	log.Println("Crypter prepared")
	ch <- crypter
}

func initMongo(ch chan<- mongoer.IMongoer) {
	mongoErrors := []error{}
	mongo, err := mongoer.Get(readConfigField("MONGO_HOST"))
	mongoErrors = append(mongoErrors, err)
	mongoErrors = append(
		mongoErrors,
		mongo.CreateCollection("user"),
	)
	mongoErrors = append(
		mongoErrors,
		mongo.CreateCollection("net"),
	)
	mongoErrors = append(
		mongoErrors,
		mongo.CreateCollection("trades"),
	)
	mongoErrors = append(
		mongoErrors,
		mongo.CreateIndex("user", "key", "hashed"),
	)
	mongoErrors = append(
		mongoErrors,
		mongo.CreateIndex("trades", "ukey", "hashed"),
	)
	mongoErrors = append(
		mongoErrors,
		mongo.CreateIndex("trades", "mkey", "hashed"),
	)
	for _, err := range mongoErrors {
		if err != nil {
			log.Panic(err)
		}
	}
	log.Println("Connected to mongo")
	ch <- mongo
}

func initAPIs(
	redis_client rediser.IRediser,
	mongo_client mongoer.IMongoer,
	crypt_client crypter.ICrypter,
) {
	m := map[string]string{
		"name":      readConfigField("MARKET_NAME"),
		"mkey":      crypt_client.Pub(),
		"descr":     readConfigField("MARKET_DESCR"),
		"img":       readConfigField("MARKET_IMG"),
		"worktime":  readConfigField("MARKET_WORKTIME"),
		"fee":       readConfigField("MARKET_FEE"),
		"delimiter": readConfigField("MARKET_DELIMITER"),
	}
	mkt_info, err := json.Marshal(m)
	if err != nil {
		log.Panic("unable to parse response")
	}
	info.Setup(mkt_info, mongo_client, crypt_client, redis_client)
	market.Setup(mkt_info, mongo_client, crypt_client, redis_client)
	operator.Setup(mkt_info, mongo_client, crypt_client, redis_client)
	user.Setup(mongo_client, crypt_client, redis_client)
	log.Println("Setup sucess...")
}

func init() {
	godotenv.Load()
	go initGraylog()
	redischan := make(chan rediser.IRediser)
	mongochan := make(chan mongoer.IMongoer)
	cryptchan := make(chan crypter.ICrypter)
	go initRedis(redischan)
	go initMongo(mongochan)
	go initCrypt(cryptchan)
	initAPIs(
		<-redischan,
		<-mongochan,
		<-cryptchan,
	)
}

func readConfigField(field string) string {
	envar, found := os.LookupEnv(field)
	if !found {
		log.Fatal("problem loading .ENV field: ", envar)
	}
	return envar
}
