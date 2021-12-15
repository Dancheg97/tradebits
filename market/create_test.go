package market

import (
	"orb/calc"
	"orb/data"
	"testing"
)

func TestCreateNewMarket(t *testing.T) {
	dummyAdress := calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	err := Create(
		dummyAdress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if err != nil {
		t.Error(err)
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(dummyAdress)
}

func TestCreateMarketBadAdress(t *testing.T) {
	var badAdress = []byte{0, 1, 2, 3}
	dummyName := string(calc.Rand()[0:16])
	err := Create(
		badAdress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if err == nil {
		t.Error("market with bad adress should not be created")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(badAdress)
}

func TestCreateMarketBadName(t *testing.T) {
	dummyAdress := calc.Rand()
	err := Create(
		dummyAdress,
		"ola",
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if err == nil {
		t.Error("market with name that small should not be created")
	}
	data.TestRM([]byte("ola"))
	data.TestRM(dummyAdress)
}

func TestCreateMarketBadDescription(t *testing.T) {
	dummyAdress := calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	err := Create(
		dummyAdress,
		dummyName,
		dummyMessageKey,
		"coca cola",
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if err == nil {
		t.Error("market with description that small should not be created")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(dummyAdress)
}

func TestCreateMarketBadFee(t *testing.T) {
	dummyAdress := calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	err := Create(
		dummyAdress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		502,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if err == nil {
		t.Error("market with fee that big should not be created")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(dummyAdress)
}

func TestCreateMarketBadWorkTime(t *testing.T) {
	dummyAdress := calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	err := Create(
		dummyAdress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		"9-21",
		dummyDelimiter,
	)
	if err == nil {
		t.Error("market with work time that small should not be created")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(dummyAdress)
}

func TestCreateMarketBadDelimited(t *testing.T) {
	dummyAdress := calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	err := Create(
		dummyAdress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		42,
	)
	if err == nil {
		t.Error("market with delimiter that big should not be created")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(dummyAdress)
}

func TestCreateUntrimmedMarket(t *testing.T) {
	var adress1 = calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	err := Create(
		adress1,
		" Name cont  spacezz  ",
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if err == nil {
		t.Error("market with same name should not be created")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(adress1)
}

func TestCreateBadDescription(t *testing.T) {
	var adress1 = calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	dummyDescription := "We will fuck your ass!! We will fuck your ass!! We will fuck your ass!! We will fuck your ass!! We will fuck your ass!! We will fuck your ass!! We will fuck your ass!! We will fuck your ass!! We will fuck your ass!! We will fuck your ass!! We will fuck your ass!! We will fuck your ass!! We will fuck your ass!!"
	err := Create(
		adress1,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	data.TestRM([]byte(dummyName))
	data.TestRM(adress1)
	if err == nil {
		t.Error("market with bad description should not be created")
	}
}

func TestCreateBadNameMarket(t *testing.T) {
	var adress1 = calc.Rand()
	dummyName := "Shitting bitch marketplace"
	err := Create(
		adress1,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if err == nil {
		t.Error("market with same name should not be created")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(adress1)
}

func TestCreateMarketBadMessageKey(t *testing.T) {
	dummyAdress := calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	err := Create(
		dummyAdress,
		dummyName,
		calc.Rand()[0:16],
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if err == nil {
		t.Error("market with mes key that small should not be created")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(dummyAdress)
}

func TestCreateExistingMarket(t *testing.T) {
	adress := calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	Create(
		adress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	err := Create(
		adress,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if err == nil {
		t.Error("new market should not be craeted")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(adress)
}

func TestCreateMarketsWithSameName(t *testing.T) {
	var adress1 = calc.Rand()
	dummyName := string(calc.Rand()[0:16])
	Create(
		adress1,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	var adress2 = calc.Rand()
	err := Create(
		adress2,
		dummyName,
		dummyMessageKey,
		dummyDescription,
		dummyImageLink,
		dummyInputFee,
		dummyOutputFee,
		dummyWorkTime,
		dummyDelimiter,
	)
	if err == nil {
		t.Error("market with same name should not be created")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(adress1)
}
