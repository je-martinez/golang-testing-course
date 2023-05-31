package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var theTests = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"404", "/fish", http.StatusNotFound},
	}

	var app Application
	routes := app.routes()

	//create a test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	//range through test data

	pathToTemplates = "./../../templates/"

	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s expected status %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}

	}
}