package data

import (
	"os"
	"testing"

	"github.com/blevesearch/bleve/v2"
)

func removeAfterTest() {
	os.RemoveAll("data")
}
func TestIsolatedCreateBleeve(t *testing.T) {
	mapping := bleve.NewIndexMapping()
	_, err := bleve.New("data/search", mapping)
	if err != nil {
		t.Error("search engine have not been created", err)
	}
	removeAfterTest()
}

func TestIsolatedAddStuff(t *testing.T) {
	mapping := bleve.NewIndexMapping()
	index, _ := bleve.New("data/search", mapping)
	addErr := index.Index("added stuff", []byte("hola"))
	if addErr != nil {
		t.Error("problem adding some stuff to directory")
	}
	removeAfterTest()
}

func TestIsolatedAddStuffAbdSearchForThatStuff(t *testing.T) {
	mapping := bleve.NewIndexMapping()
	index, _ := bleve.New("data/search", mapping)
	index.Index("added stuff", "stuff")
	index.Index("added stuff", "added")
	query := bleve.NewQueryStringQuery("added")
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, searchErr := index.Search(searchRequest)
	if searchErr != nil {
		t.Error("search error occured")
	}
	if searchResult.Total != 1 {
		t.Error("it should find 1 result")
		t.Error("search found amount of results: ", searchResult.Total)
	}
	removeAfterTest()
}

func TestIsolatedAddAndSearch(t *testing.T) {
	adress := []byte{0, 1, 2, 3}
	name := "marketSuffHola"
	SearchAdd(name, adress)
	rez := Search(name)
	t.Error(rez)
	removeAfterTest()
}
