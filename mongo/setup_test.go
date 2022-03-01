package mongo

import "testing"

func TestOpenMongoSuccess(t *testing.T) {
	Setup("mongodb://localhost:27017")
}

func TestOpenMongoFail(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Setup("")
}

func TestCreateCollectionSuccess(t *testing.T) {

}

func TestCreateCollectionFail(t *testing.T) {

}

func TestCreateIndexSuccess(t *testing.T) {

}

func TestCreateIndexFail(t *testing.T) {

}
