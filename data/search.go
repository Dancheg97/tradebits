package data

import "github.com/blevesearch/bleve/v2"

/*
	Purpose of search engine is to give user an ability to find relevant to
request markets.
    Search engine is gonna work by following way - when some new market is
created, it's adress will be saved to maping according to new market name.
	If there are two markets craeted with the same name, both adresses are
gonna be saved to same.
*/

var searcher, _ = bleve.Open("data/search")

func SearchAdd(name string, adress []byte) {
	adressAsString := string(adress)
	searcher.Index(adressAsString, name)
}

func Search(info string) [][]byte {
	query := bleve.NewMatchQuery(info)
	search := bleve.NewSearchRequest(query)
	searchRez, _ := searcher.Search(search)
	rezArr := [][]byte{}
	for idx, hit := range searchRez.Hits {
		rezArr = append(rezArr, []byte(hit.ID))
		if idx == 30 {
			break
		}
	}
	return rezArr
}
