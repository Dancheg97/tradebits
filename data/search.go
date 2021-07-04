package data

import (
	"os"

	"github.com/blevesearch/bleve/v2"
)

/*
	Purpose of search engine is to give user an ability to find relevant to
request markets.
    Search engine is gonna work by following way - when some new market is
created, it's adress will be saved to maping according to new market name.
	If there are two markets craeted with the same name, both adresses are
gonna be saved to same.
*/

var searchPath = "/Users/danilafominyh/Documents/sync_tree_server/data/search"
var searcher = openSearch()

func openSearch() bleve.Index {
	_, existErr := os.Stat(searchPath)
	if existErr != nil {
		mapping := bleve.NewIndexMapping()
		searcher, _ := bleve.New(searchPath, mapping)
		return searcher
	}
	searcher, _ := bleve.Open(searchPath)
	return searcher
}

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
