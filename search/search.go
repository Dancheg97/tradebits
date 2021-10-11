package search

import (
	"github.com/blevesearch/bleve/v2"
)

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
