package main

import (
	"os"
	"testing"
)

var app Application

func TestMain(m *testing.M) {
	app.Session = getSession()
	os.Exit(m.Run())
}
