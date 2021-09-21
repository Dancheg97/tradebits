package search

import (
	"os"
	"path"
	"runtime"

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

var searcher = openSearch()

func openSearch() bleve.Index {
	_, filename, _, _ := runtime.Caller(0)
	searchPath := path.Dir(filename) + "/data.bleve"
	_, existErr := os.Stat(searchPath)
	if existErr != nil {
		mapping := bleve.NewIndexMapping()
		searcher, _ := bleve.New(searchPath, mapping)
		return searcher
	}
	searcher, _ := bleve.Open(searchPath)
	return searcher
}

func Add(name string, adress []byte) {
	adressAsString := string(adress)
	searcher.Index(adressAsString, name+" "+adressAsString)
}

func Change(newName string, adress []byte) {
	// отгрузить предыдущий результат
	adressAsString := string(adress)
	searcher.Delete(adressAsString)
	searcher.Index(adressAsString, newName+" "+adressAsString)
}

func Search(info string) [][]byte {
	query := bleve.NewMatchQuery(info)
	search := bleve.NewSearchRequest(query)
	search.Size = 30
	searchRez, _ := searcher.Search(search)
	rezArr := [][]byte{}
	for _, hit := range searchRez.Hits {
		rezArr = append(rezArr, []byte(hit.ID))
	}
	return rezArr
}
