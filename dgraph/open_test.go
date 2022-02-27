package dgraph

import (
	"testing"
)

func TestOpenDgraphSuccess(t *testing.T) {
	dgraphClient := newClient("localhost:9080")
	if dgraphClient == nil {
		t.Error("Error opening dgraph instance")
	}
	dgraphClient.NewTxn()
}

func TestOpenDgraphError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	newClient("")
}
