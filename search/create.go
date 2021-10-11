package search

import (
	"os"
	"path"
	"runtime"

	"github.com/blevesearch/bleve/v2"
)

var searcher = openSearch()

func openSearch() bleve.Index {
	_, filename, _, _ := runtime.Caller(0)
	searchPath := path.Dir(filename) + "/bleve"
	_, existErr := os.Stat(searchPath)
	if existErr != nil {
		mapping := bleve.NewIndexMapping()
		searcher, _ := bleve.New(searchPath, mapping)
		return searcher
	}
	searcher, _ := bleve.Open(searchPath)
	return searcher
}
