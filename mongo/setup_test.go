package mongo

import "testing"

func TestSetupSucess(t *testing.T) {
	Setup("mongodb://localhost:27017")
}

func TestSetupFail(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Setup("")
}
