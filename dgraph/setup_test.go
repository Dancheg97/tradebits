package dgraph

import (
	"testing"
)

func TestSetupSuccess(t *testing.T) {
	Setup("localhost:9080", "data.gql")
}

func TestSetupError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Setup("", "data.gql")
}
