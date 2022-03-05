package api

import (
	"net/http"
	"testing"
)


func TestInfoMarketGet(t *testing.T) {
	mockWriter := MockHttpWriter{}
	InfoMarketGet(mockWriter, &http.Request{})
	if mockWriter.expectedStatus != 200 {
		t.Error("wrong status code in respone")
	}
}
