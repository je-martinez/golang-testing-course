package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_app_authenticate(t *testing.T) {
	var tests = []struct {
		name               string
		requestBody        string
		expectedStatusCode int
	}{
		{"valid user", `{"email":"admin@example.com","password":"secret"}`, http.StatusOK},
		{"not json", `invalid json`, http.StatusBadRequest},
		{"empty json", `{}`, http.StatusUnauthorized},
		{"empty email", `{"email":""}`, http.StatusUnauthorized},
		{"empty password", `{"email":"admin@example.com","password":""}`, http.StatusUnauthorized},
		{"invalid user", `{"email":"admin@someotherdomain.com","password":"secret"}`, http.StatusUnauthorized},
	}

	for _, e := range tests {
		var reader io.Reader
		reader = strings.NewReader(e.requestBody)
		req, _ := http.NewRequest("POST", "/auth", reader)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.authenticate)
		handler.ServeHTTP(rr, req)

		if e.expectedStatusCode != rr.Code {
			t.Errorf("%s: returned wrong status code: expected %d but got %d", e.name, e.expectedStatusCode, rr.Code)
		}
	}
}
