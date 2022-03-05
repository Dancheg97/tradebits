package api

import "net/http"

type MockHttpWriter struct {
	expectedStatus int
}

func (m MockHttpWriter) Header() http.Header {
	return http.Header{}
}

func (m MockHttpWriter) Write(bytes []byte) (int, error) {
	return 0, nil
}

func (m MockHttpWriter) WriteHeader(statusCode int) {
	m.expectedStatus = statusCode
}
