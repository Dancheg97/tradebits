Сейчас:
/ - Сделать удобное API для взаимодействия с базой данных (для манипуляции различными объекетами базы данных)

В перспективе:
/ - сделать для базы данных скан выделенных дисков
/ - сделать для базы данных автоскейлинг в рамках выделенных дисков
/ - сделать для базы данных автоскейлинг при условии замедления

OLD_CODE:

func getDataBases() map[byte]*leveldb.DB {
	var dataBases = make(map[byte]*leveldb.DB)
	for i := 0; i < 256; i++ {
		index := fmt.Sprint(uint8(i))
		var db, _ = leveldb.OpenFile("database/storage/"+index, nil)
		dataBases[byte(i)] = db
	}
	return dataBases
}

var bases = getDataBases()

func WriteValue(key []byte, value []byte) {
	dbKey := key[0]
	bases[dbKey].Put(key, value, nil)
}

func ReadValue(key []byte) []byte {
	dbKey := key[0]
	value, readError := bases[dbKey].Get(key, nil)
	if readError != nil {
		return nil
	}
	return value
}

func ReadBalance(adress []byte) uint64 {
	var byteBalance = ReadValue(adress)
	if byteBalance == nil {
		return 0
	}
	return binary.LittleEndian.Uint64(byteBalance)
}

func WriteBalance(adress []byte, balance uint64) {
	balanceAsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(balanceAsBytes, balance)
	WriteValue(adress, balanceAsBytes)
}
