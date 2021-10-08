package market

var dummyMessageKey = []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2}
var dummyImageLink = "test.imagelink/thereisnoimagebythislink"
var dummyDescription = `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrbled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets ctaining Lorem Ipsum pages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`
var dummyInputFee = uint64(100)
var dummyOutputFee = uint64(100)
var dummyWorkTime = "+3GMT 9:00 - 21:00"
var dummyDelimiter = uint64(2)

// func TestAttachUnboundedTrades(t *testing.T) {
// 	var adress = calc.Rand()
// 	dummyName := string(calc.Rand()[0:16])
// 	Create(
// 		adress,
// 		dummyName,
// 		dummyMessageKey,
// 		dummyDescription,
// 		dummyImageLink,
// 		dummyInputFee,
// 		dummyOutputFee,
// 		dummyWorkTime,
// 		dummyDelimiter,
// 	)
// 	mkt := Get(adress)
// 	sell := trade.Sell{
// 		Offer:   100,
// 		Recieve: 100,
// 	}
// 	buy := trade.Buy{
// 		Offer:   100,
// 		Recieve: 100,
// 	}
// 	buyAttached := mkt.AttachBuy(&buy)
// 	sellAttached := mkt.AttachSell(&sell)
// 	if buyAttached || sellAttached {
// 		t.Error("those trades should not be attached cuz they are unbounded")
// 	}
// 	data.TestRM([]byte(dummyName))
// 	data.TestRM(adress)
// }

// func TestAttachToLookedMarket(t *testing.T) {
// 	var adress = calc.Rand()
// 	dummyName := string(calc.Rand()[0:16])
// 	Create(
// 		adress,
// 		dummyName,
// 		dummyMessageKey,
// 		dummyDescription,
// 		dummyImageLink,
// 		dummyInputFee,
// 		dummyOutputFee,
// 		dummyWorkTime,
// 		dummyDelimiter,
// 	)
// 	mkt := Look(adress)
// 	sell := trade.Sell{
// 		Offer:   100,
// 		Recieve: 100,
// 	}
// 	buy := trade.Buy{
// 		Offer:   100,
// 		Recieve: 100,
// 	}
// 	buyAttached := mkt.AttachBuy(&buy)
// 	sellAttached := mkt.AttachSell(&sell)
// 	if buyAttached || sellAttached {
// 		t.Error("those trades should not be attached cuz they are unbounded")
// 	}
// 	data.TestRM([]byte(dummyName))
// 	data.TestRM(adress)
// }

// // func TestAttachSingleNormalBuy(t *testing.T) {
// // 	var marketAdress = calc.Rand()
// // 	dummyName := string(calc.Rand()[0:16])
// // 	Create(
// // 		marketAdress,
// // 		dummyName,
// // 		dummyMessageKey,
// // 		dummyDescription,
// // 		dummyImageLink,
// // 		dummyInputFee,
// // 		dummyOutputFee,
// // 		dummyWorkTime,
// // 		dummyDelimiter,
// // 	)
// // 	mkt := Get(marketAdress)
// // 	var userAdress = calc.Rand()
// // 	dummyUserName := string(calc.Rand()[0:8])
// // 	user.Create(
// // 		userAdress,
// // 		dummyMessageKey,
// // 		dummyUserName,
// // 	)
// // 	usr := user.Get(userAdress)
// // 	usr.Balance = 100
// // 	buy := trade.Buy{
// // 		Offer:   100,
// // 		Recieve: 100,
// // 	}
// // 	attachedToUser := usr.AttachBuy(&buy)
// // 	if !attachedToUser {
// // 		t.Error("trade should be attached to user")
// // 	}
// // 	attachedToMarket := mkt.AttachBuy(&buy)
// // 	if !attachedToMarket {
// // 		t.Error("trade should be attached to market")
// // 	}
// // 	data.TestRM([]byte(dummyUserName))
// // 	data.TestRM([]byte(dummyName))
// // 	data.TestRM(marketAdress)
// // 	data.TestRM(userAdress)
// // }

// // func TestAttachSingleNormalSell(t *testing.T) {
// // 	var marketAdress = calc.Rand()
// // 	dummyName := string(calc.Rand()[0:16])
// // 	Create(
// // 		marketAdress,
// // 		dummyName,
// // 		dummyMessageKey,
// // 		dummyDescription,
// // 		dummyImageLink,
// // 		dummyInputFee,
// // 		dummyOutputFee,
// // 		dummyWorkTime,
// // 		dummyDelimiter,
// // 	)
// // 	mkt := Get(marketAdress)
// // 	var userAdress = calc.Rand()
// // 	dummyUserName := string(calc.Rand()[0:8])
// // 	user.Create(
// // 		userAdress,
// // 		dummyMessageKey,
// // 		dummyUserName,
// // 	)
// // 	usr := user.Get(userAdress)
// // 	usr.Balances[string(marketAdress)] = 100
// // 	buy := trade.Sell{
// // 		Offer:   100,
// // 		Recieve: 100,
// // 	}
// // 	attachedToUser := usr.AttachSell(&buy, marketAdress)
// // 	if !attachedToUser {
// // 		t.Error("trade should be attached to user")
// // 	}
// // 	attachedToMarket := mkt.AttachSell(&buy)
// // 	if !attachedToMarket {
// // 		t.Error("trade should be attached to market")
// // 	}
// // 	data.TestRM([]byte(dummyUserName))
// // 	data.TestRM([]byte(dummyName))
// // 	data.TestRM(marketAdress)
// // 	data.TestRM(userAdress)
// // }

// func TestTwoUserTradesWithSameOffers(t *testing.T) {
// 	var marketAdress = calc.Rand()
// 	dummyName := string(calc.Rand()[0:16])
// 	Create(
// 		marketAdress,
// 		dummyName,
// 		dummyMessageKey,
// 		dummyDescription,
// 		dummyImageLink,
// 		dummyInputFee,
// 		dummyOutputFee,
// 		dummyWorkTime,
// 		dummyDelimiter,
// 	)
// 	mkt := Get(marketAdress)
// 	var firstUserAdress = calc.Rand()
// 	dummyUserName := string(calc.Rand()[0:8])
// 	user.Create(
// 		firstUserAdress,
// 		dummyMessageKey,
// 		dummyUserName,
// 	)
// 	firstUser := user.Get(firstUserAdress)
// 	firstUser.Balance = 100
// 	buy := trade.Buy{
// 		Offer:   100,
// 		Recieve: 100,
// 	}
// 	firstUser.AttachBuy(&buy)
// 	mkt.AttachBuy(&buy)
// 	firstUser.Save()
// 	var secondUserAdress = calc.Rand()
// 	dummyUserName2 := string(calc.Rand()[0:8])
// 	user.Create(
// 		secondUserAdress,
// 		dummyMessageKey,
// 		dummyUserName2,
// 	)
// 	secondUser := user.Get(secondUserAdress)
// 	secondUser.Balances[string(marketAdress)] = 100
// 	sell := trade.Sell{
// 		Offer:   100,
// 		Recieve: 100,
// 	}
// 	secondUser.AttachSell(&sell, marketAdress)
// 	mkt.AttachSell(&sell)
// 	secondUser.Save()
// 	time.Sleep(time.Second * 1)
// 	usr1check := user.Look(firstUserAdress)
// 	if usr1check.Balance != 0 {
// 		t.Error("first user main balance fshould be equal to zero")
// 	}
// 	if usr1check.Balances[string(marketAdress)] != 100 {
// 		t.Error("first user market balance should be equal to 100")
// 	}
// 	usr2check := user.Look(secondUserAdress)
// 	if usr2check.Balances[string(marketAdress)] != 0 {
// 		t.Error("market balance of second user should be equal to zero")
// 	}
// 	if usr2check.Balance != 100 {
// 		t.Error("second user main balance should be equal to 100")
// 	}
// 	data.TestRM([]byte(dummyName))
// 	data.TestRM([]byte(dummyUserName))
// 	data.TestRM([]byte(dummyUserName2))
// 	data.TestRM(firstUserAdress)
// 	data.TestRM(secondUserAdress)
// 	data.TestRM(marketAdress)
// }

// func TestFourUserTradesWithRandomOffers(t *testing.T) {
// 	var marketAdress = calc.Rand()
// 	dummyName := string(calc.Rand()[0:16])
// 	Create(
// 		marketAdress,
// 		dummyName,
// 		dummyMessageKey,
// 		dummyDescription,
// 		dummyImageLink,
// 		dummyInputFee,
// 		dummyOutputFee,
// 		dummyWorkTime,
// 		dummyDelimiter,
// 	)
// 	mkt := Get(marketAdress)
// 	var firstUserAdress = calc.Rand()
// 	dummyUserName := string(calc.Rand()[0:8])
// 	user.Create(
// 		firstUserAdress,
// 		dummyMessageKey,
// 		dummyUserName,
// 	)
// 	firstUser := user.Get(firstUserAdress)
// 	firstUser.Balance = 300
// 	firstUserTrade := trade.Buy{
// 		Offer:   270,
// 		Recieve: 130,
// 	}
// 	firstUser.AttachBuy(&firstUserTrade)
// 	mkt.AttachBuy(&firstUserTrade)
// 	firstUser.Save()
// 	var secondUserAdress = calc.Rand()
// 	dummyUserName2 := string(calc.Rand()[0:8])
// 	user.Create(
// 		secondUserAdress,
// 		dummyMessageKey,
// 		dummyUserName2,
// 	)
// 	secondUser := user.Get(secondUserAdress)
// 	secondUser.Balances[string(marketAdress)] = 150
// 	secondUserTrade := trade.Sell{
// 		Offer:   80,
// 		Recieve: 130,
// 	}
// 	secondUser.AttachSell(&secondUserTrade, marketAdress)
// 	mkt.AttachSell(&secondUserTrade)
// 	secondUser.Save()
// 	var thirdUserAdress = calc.Rand()
// 	dummyUserName3 := string(calc.Rand()[0:8])
// 	user.Create(
// 		thirdUserAdress,
// 		dummyMessageKey,
// 		dummyUserName3,
// 	)
// 	thirdUser := user.Get(thirdUserAdress)
// 	thirdUser.Balances[string(marketAdress)] = 150
// 	thirdUserTrade := trade.Sell{
// 		Offer:   20,
// 		Recieve: 15,
// 	}
// 	thirdUser.AttachSell(&thirdUserTrade, marketAdress)
// 	mkt.AttachSell(&thirdUserTrade)
// 	thirdUser.Save()
// 	time.Sleep(time.Second * 1)
// 	firstUserCheck := user.Look(firstUserAdress)
// 	if firstUserCheck.Balance != 30 {
// 		t.Error("first user balance should be equal to 30")
// 	}
// 	if firstUserCheck.Balances[string(marketAdress)] != 100 {
// 		t.Error("first user market balance should be equal to 100")
// 	}
// 	secondUserCheck := user.Look(secondUserAdress)
// 	if secondUserCheck.Balance != 130 {
// 		t.Error("second user balance should be equal to 130")
// 	}
// 	if secondUserCheck.Balances[string(marketAdress)] != 70 {
// 		t.Error("second user market balance should be equal to 70")
// 	}
// 	if len(mkt.Pool.Buys) != 1 {
// 		t.Error("market pool length should be equal to one")
// 	}
// 	if mkt.Pool.Buys[0].Offer != 125 {
// 		t.Error("current offer of market buy should be equal to 125")
// 		t.Error(mkt.Pool.Buys[0].Offer)
// 	}
// 	if mkt.Pool.Buys[0].Recieve != 30 {
// 		t.Error("current offer of market buy should be equal to 30")
// 		t.Error(mkt.Pool.Buys[0].Recieve)
// 	}
// 	if len(mkt.Pool.Sells) != 0 {
// 		t.Error("there should not be any active market sell")
// 	}
// 	if len(mkt.Pool.Outputs) != 0 {
// 		t.Error("there should not be any market outputs")
// 	}
// 	data.TestRM(marketAdress)
// 	data.TestRM(firstUserAdress)
// 	data.TestRM(secondUserAdress)
// 	data.TestRM(thirdUserAdress)
// 	data.TestRM([]byte(dummyName))
// 	data.TestRM([]byte(dummyUserName))
// 	data.TestRM([]byte(dummyUserName2))
// 	data.TestRM([]byte(dummyUserName3))
// }

// func TestAttachFirstlySellThanBuy(t *testing.T) {
// 	var marketAdress = calc.Rand()
// 	dummyName := string(calc.Rand()[0:16])
// 	Create(
// 		marketAdress,
// 		dummyName,
// 		dummyMessageKey,
// 		dummyDescription,
// 		dummyImageLink,
// 		dummyInputFee,
// 		dummyOutputFee,
// 		dummyWorkTime,
// 		dummyDelimiter,
// 	)
// 	mkt := Get(marketAdress)
// 	var firstUserAdress = calc.Rand()
// 	dummyUserName := string(calc.Rand()[0:8])
// 	user.Create(
// 		firstUserAdress,
// 		dummyMessageKey,
// 		dummyUserName,
// 	)
// 	firstUser := user.Get(firstUserAdress)
// 	firstUser.Balance = 300
// 	firstUserTrade := trade.Buy{
// 		Offer:   270,
// 		Recieve: 130,
// 	}
// 	firstUser.AttachBuy(&firstUserTrade)
// 	firstUser.Save()
// 	var secondUserAdress = calc.Rand()
// 	dummyUserName2 := string(calc.Rand()[0:8])
// 	user.Create(
// 		secondUserAdress,
// 		dummyMessageKey,
// 		dummyUserName2,
// 	)
// 	secondUser := user.Get(secondUserAdress)
// 	secondUser.Balances[string(marketAdress)] = 150
// 	secondUserTrade := trade.Sell{
// 		Offer:   80,
// 		Recieve: 130,
// 	}
// 	secondUser.AttachSell(&secondUserTrade, marketAdress)
// 	secondUser.Save()
// 	mkt.AttachSell(&secondUserTrade)
// 	mkt.AttachBuy(&firstUserTrade)
// 	time.Sleep(time.Second)
// 	firstUserCheck := user.Look(firstUserAdress)
// 	secondUserCheck := user.Look(secondUserAdress)
// 	if firstUserCheck.Balance != 30 {
// 		t.Error("first user balance should be equal to 30")
// 	}
// 	if firstUserCheck.Balances[string(marketAdress)] != 80 {
// 		t.Error("first user market balance should be equal to 80")
// 	}
// 	if secondUserCheck.Balance != 130 {
// 		t.Error("second user main balance should be equal to 130")
// 	}
// 	if secondUserCheck.Balances[string(marketAdress)] != 70 {
// 		t.Error("second user market balance should be equal to 70")
// 	}
// 	if mkt.Pool.Buys[0].Offer != 140 {
// 		t.Error("active buy offer should be equal tp 140")
// 	}
// 	if mkt.Pool.Buys[0].Recieve != 50 {
// 		t.Error("acrive buy recieve should be equal to 50")
// 	}
// 	data.TestRM(marketAdress)
// 	data.TestRM(firstUserAdress)
// 	data.TestRM(secondUserAdress)
// 	data.TestRM([]byte(dummyName))
// 	data.TestRM([]byte(dummyUserName))
// 	data.TestRM([]byte(dummyUserName2))
// }

// func TestIfUserHasTrdadesWhenHeHaveSome(t *testing.T) {
// 	var marketAdress = calc.Rand()
// 	dummyName := string(calc.Rand()[0:16])
// 	Create(
// 		marketAdress,
// 		dummyName,
// 		dummyMessageKey,
// 		dummyDescription,
// 		dummyImageLink,
// 		dummyInputFee,
// 		dummyOutputFee,
// 		dummyWorkTime,
// 		dummyDelimiter,
// 	)
// 	mkt := Get(marketAdress)
// 	sell := trade.Sell{
// 		Adress:  []byte{0},
// 		Offer:   1,
// 		Recieve: 5,
// 	}
// 	mkt.AttachSell(&sell)
// 	userHaveSell := mkt.HasTrades([]byte{0})
// 	if !userHaveSell {
// 		t.Error("error, this user should have some sell trade")
// 	}
// 	buy := trade.Buy{
// 		Adress:  []byte{0},
// 		Offer:   1,
// 		Recieve: 5,
// 	}
// 	mkt.AttachBuy(&buy)
// 	userHaveBuy := mkt.HasTrades([]byte{0})
// 	if !userHaveBuy {
// 		t.Error("error, this user should have some buy trade")
// 	}
// 	data.TestRM(marketAdress)
// 	data.TestRM([]byte(dummyName))
// }

// func TestIfUserHasTradesWhenHeDont(t *testing.T) {
// 	var marketAdress = calc.Rand()
// 	dummyName := string(calc.Rand()[0:16])
// 	Create(
// 		marketAdress,
// 		dummyName,
// 		dummyMessageKey,
// 		dummyDescription,
// 		dummyImageLink,
// 		dummyInputFee,
// 		dummyOutputFee,
// 		dummyWorkTime,
// 		dummyDelimiter,
// 	)
// 	mkt := Get(marketAdress)

// 	userHaveSomeTrades := mkt.HasTrades([]byte{1})
// 	if userHaveSomeTrades {
// 		t.Error("there should not be any active trades for that adress")
// 	}

// 	data.TestRM(marketAdress)
// 	data.TestRM([]byte(dummyName))
// }

// func TestUserCancelBuy(t *testing.T) {
// 	var marketAdress = calc.Rand()
// 	dummyName := string(calc.Rand()[0:16])
// 	Create(
// 		marketAdress,
// 		dummyName,
// 		dummyMessageKey,
// 		dummyDescription,
// 		dummyImageLink,
// 		dummyInputFee,
// 		dummyOutputFee,
// 		dummyWorkTime,
// 		dummyDelimiter,
// 	)
// 	mkt := Get(marketAdress)
// 	var userAdress = calc.Rand()
// 	dummyUserName := string(calc.Rand()[0:8])
// 	user.Create(
// 		userAdress,
// 		dummyMessageKey,
// 		dummyUserName,
// 	)
// 	usr := user.Get(userAdress)
// 	usr.Balance = 300
// 	usrBuy := trade.Buy{
// 		Offer:   270,
// 		Recieve: 130,
// 	}
// 	usr.AttachBuy(&usrBuy)
// 	mkt.AttachBuy(&usrBuy)
// 	usr.Save()
// 	mkt.CancelTrades(userAdress)
// 	time.Sleep(time.Second)
// 	userCheck := user.Look(userAdress)
// 	if len(mkt.Pool.Buys) != 0 {
// 		t.Error("the trade have not been cancelled, there should not be active trades on the market")
// 	}
// 	if userCheck.Balance != 300 {
// 		t.Error("user main balance should be equal to 300, cuz his trade has benn cancelled")
// 	}
// 	data.TestRM(marketAdress)
// 	data.TestRM(userAdress)
// 	data.TestRM([]byte(dummyName))
// 	data.TestRM([]byte(dummyUserName))
// }

// // func TestUserCancelSell(t *testing.T) {
// // 	var marketAdress = calc.Rand()
// // 	dummyName := string(calc.Rand()[0:16])
// // 	Create(
// // 		marketAdress,
// // 		dummyName,
// // 		dummyMessageKey,
// // 		dummyDescription,
// // 		dummyImageLink,
// // 		dummyInputFee,
// // 		dummyOutputFee,
// // 		dummyWorkTime,
// // 		dummyDelimiter,
// // 	)
// // 	mkt := Get(marketAdress)
// // 	var userAdress = calc.Rand()
// // 	dummyUserName := string(calc.Rand()[0:8])
// // 	user.Create(
// // 		userAdress,
// // 		dummyMessageKey,
// // 		dummyUserName,
// // 	)
// // 	usr := user.Get(userAdress)
// // 	usr.Balances[string(marketAdress)] = 150
// // 	sell := trade.Sell{
// // 		Offer:   80,
// // 		Recieve: 130,
// // 	}
// // 	usr.AttachSell(&sell, marketAdress)
// // 	usr.Save()
// // 	mkt.AttachSell(&sell)
// // 	mkt.CancelTrades(userAdress)
// // 	time.Sleep(time.Second)
// // 	userCheck := user.Look(userAdress)
// // 	if len(mkt.Pool.Sells) != 0 {
// // 		t.Error("there should not be any active sell, because user cancelled that")
// // 	}
// // 	if userCheck.Balances[string(marketAdress)] != 150 {
// // 		t.Error("user market balance should be equal to 150 after sell trade cancellation")
// // 	}
// // 	data.TestRM(marketAdress)
// // 	data.TestRM(userAdress)
// // 	data.TestRM([]byte(dummyName))
// // 	data.TestRM([]byte(dummyUserName))
// // }
