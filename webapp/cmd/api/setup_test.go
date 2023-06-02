package main

import (
	"os"
	"testing"
	"webapp/pkg/repository/dbrepo"
)

var app Application

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "very-secret"
	os.Exit(m.Run())
}
