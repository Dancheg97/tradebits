package mongo

import "testing"

func TestOpenMongoSuccess(t *testing.T) {
	openMongo("mongodb://localhost:27017")
}

func TestOpenMongoFail(t *testing.T) {
	
}

func TestCreateCollectionSuccess(t *testing.T) {
	Setup("mongodb://localhost:27017")
}

func TestCreateCollectionFail(t *testing.T) {

}

func TestCreateIndexSuccess(t *testing.T) {

}

func TestCreateIndexFail(t *testing.T) {

}
