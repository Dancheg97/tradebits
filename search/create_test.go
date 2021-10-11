package search

import (
	"os"
	"path"
	"runtime"
	"testing"
	"time"
)

func TestRecreateSearcher(t *testing.T) {
	time.Sleep(time.Millisecond*620)
	searcher.Close()
	_, filename, _, _ := runtime.Caller(0)
	searchPath := path.Dir(filename) + "/bleve"
	os.RemoveAll(searchPath)
	searcher = openSearch()
	time.Sleep(time.Millisecond*620)
}
