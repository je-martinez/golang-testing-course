package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_application_routes(t *testing.T) {

	var registered = []struct {
		route  string
		method string
	}{
		{"/", "GET"},
		{"/login", "POST"},
		{"/static/*", "GET"},
	}

	var app Application
	mux := app.routes()

	chiRoutes := mux.(chi.Routes)

	for _, route := range registered {
		if !routeExist(route.route, route.method, chiRoutes) {
			t.Errorf("route %s is not registered", route.route)
		}
		//Check to see if the route exist
	}
}

func routeExist(testRoute string, testMethod string, chiRoutes chi.Routes) bool {
	found := false
	_ = chi.Walk(chiRoutes, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
			found = true
		}
		return nil
	})
	return found
}