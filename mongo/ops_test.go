package mongo

import "testing"

func TestCheck(t *testing.T) {
	OpenMongo("mongodb://localhost:27017")
	CreateCollection("test")
	CreateIndex("key", "")
}
