package data

import (
	"os"
	"testing"

	"github.com/blevesearch/bleve/v2"
)

func TestCreateBleeve(t *testing.T) {
	mapping := bleve.NewIndexMapping()
	_, err := bleve.New("search", mapping)
	if err != nil {
		t.Error("index db have not been created")
	}
	os.Remove("data")
}

func TestAddStuff(t *testing.T) {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("search", mapping)
	if err != nil {
		t.Error("index db have not been created")
	}
	index.Index("added stuff" , []byte("hola"))
	os.Remove("search")
}