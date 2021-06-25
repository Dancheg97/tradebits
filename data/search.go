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

func SearchAdd(info string, adress []byte) {
	searcher.Index(info, adress)
}

func Search(info string) string {
	query := bleve.NewMatchQuery("text")
	search := bleve.NewSearchRequest(query)
	searchResults, _ := searcher.Search(search)
	rez := searchResults.String()
	return rez
}
