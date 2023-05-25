package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_addIpToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAdd    bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}

	var app Application

	//create a dummy handler

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//make sure that value exists in the context
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not present")
		}

		//make sure we got a stirng back
		ip, ok := val.(string)
		if !ok {
			t.Error("not string")
		}
		t.Log(ip)
	})

	for _, e := range tests {
		//create the handler to test
		handlerToTest := app.addIpToContext(nextHandler)
		req := httptest.NewRequest("GET", "http://testing", nil)

		if e.emptyAdd {
			req.RemoteAddr = ""
		}

		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFromContext(t *testing.T) {
	//create app variable
	var app Application

	//get a context
	var ctx = context.Background()

	//put something in the context
	ctx = context.WithValue(ctx, contextUserKey, "192.168.1.1")

	//call the function
	ip := app.ipFromContext(ctx)

	//perform the test
	if !strings.EqualFold("192.168.1.1", ip) {
		t.Error("wrong value returned from context")
	}
}
