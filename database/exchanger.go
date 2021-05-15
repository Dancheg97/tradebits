package database

type Exchanger struct {
	Name               string `json:"Name"`
	Image              []byte `json:"Image"`
	Pledge             uint64 `json:"Pledge"`
	MessageKey         []byte `json:"MessageKey"`
	GoodFeedBacksCount uint64 `json:"GoodFeedBacksCount"`
	BadFeedBacksCount  uint64 `json:"BadFeedBacksCount"`
	RequestsLink       []byte `json:"RequestsLink"`
}

func NewExchanger(adress []byte) {

}
