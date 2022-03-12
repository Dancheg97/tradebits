package connect

import (
	"encoding/json"
	"net/http"
	"tradebits/mongoer"
)

type InfoResponse struct {
	Name string `json:"name"`
	Mkey string `json:"mkey"`
}

func saveInformation(adress string, mongo mongoer.IMongoer) error {
	response, err := http.Get("http://" + adress + "/info/market")
	if err != nil {
		return err
	}
	dec := json.NewDecoder(response.Body)
	resp := InfoResponse{}
	err = dec.Decode(&resp)
	if err != nil {
		return err
	}
	newMarket := map[string]string{
		"name": resp.Name,
		"mkey": resp.Mkey,
		"link": adress,
	}
	err = mongo.Put("net", newMarket)
	if err != nil {
		return err
	}
	return nil
}
