package search

import (
	"os"
	"path"
	"runtime"
	"testing"
	"time"
)

func TestRecreateSearcher(t *testing.T) {
	time.Sleep(time.Second * 8)
	searcher.Close()
	_, filename, _, _ := runtime.Caller(0)
	searchPath := path.Dir(filename) + "/bleve"
	os.RemoveAll(searchPath)
	openSearch()
}
