package data

import (
	"reflect"
	"testing"
)

// func TestIsolatedCreateBleeve(t *testing.T) {
// 	mapping := bleve.NewIndexMapping()
// 	_, err := bleve.New("data/search", mapping)
// 	if err != nil {
// 		t.Error("search engine have not been created", err)
// 	}

// }

// func TestIsolatedAddStuff(t *testing.T) {
// 	mapping := bleve.NewIndexMapping()
// 	searcher, _ := bleve.New("data/search", mapping)
// 	addErr := searcher.Index("added stuff", []byte("hola"))
// 	if addErr != nil {
// 		t.Error("problem adding some stuff to directory")
// 	}

// }

// func TestIsolatedAddStuffAbdSearchForThatStuff(t *testing.T) {
// 	mapping := bleve.NewIndexMapping()
// 	searcher, _ := bleve.New("data/search", mapping)
// 	searcher.Index("added stuff", "stuff")
// 	searcher.Index("added stuff", "added")
// 	query := bleve.NewQueryStringQuery("added")
// 	searchRequest := bleve.NewSearchRequest(query)
// 	searchResult, searchErr := searcher.Search(searchRequest)
// 	if searchErr != nil {
// 		t.Error("search error occured")
// 	}
// 	if searchResult.Total != 1 {
// 		t.Error("it should find 1 result")
// 		t.Error("search found amount of results: ", searchResult.Total)
// 	}

// }

// func TestIsolatedAddAndSearch(t *testing.T) {
// 	mapping := bleve.NewIndexMapping()
// 	searcher, _ := bleve.New("data/search", mapping)
// 	name := "holla"
// 	adressAsString := string([]byte{0, 1, 2, 3})
// 	searcher.Index(adressAsString, name)
// 	query := bleve.NewMatchQuery(name)
// 	search := bleve.NewSearchRequest(query)
// 	searchResults, _ := searcher.Search(search)
// 	hit := searchResults.Hits[0]
// 	if !reflect.DeepEqual([]byte(hit.ID), []byte{0, 1, 2, 3}) {
// 		t.Error("search result didn't match")
// 	}

// }

// func TestIsolatedAddMultipleValuesToSameString(t *testing.T) {
// 	mapping := bleve.NewIndexMapping()
// 	searcher, _ := bleve.New("data/search", mapping)
// 	searcher.Index("hola", "hola")
// 	err := searcher.Index("hola", "holasasd")
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestIsolatedAddMultipleValuesToSameIndexAndGetThem(t *testing.T) {
// 	mapping := bleve.NewIndexMapping()
// 	searcher, _ := bleve.New("data/search", mapping)
// 	searcher.Index(string([]byte{0, 1, 2, 3}), "hola")
// 	searcher.Index(string([]byte{0, 1, 2, 4}), "hola")
// 	query := bleve.NewMatchQuery("hola")
// 	search := bleve.NewSearchRequest(query)
// 	searchResults, _ := searcher.Search(search)
// 	for _, hit := range searchResults.Hits {
// 		t.Error([]byte(hit.ID))
// 	}

// }
