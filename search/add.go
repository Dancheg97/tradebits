package search

func Add(name string, adress []byte) {
	adressAsString := string(adress)
	searcher.Index(adressAsString, name+" "+adressAsString)
}
