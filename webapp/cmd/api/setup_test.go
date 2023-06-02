package main

import (
	"os"
	"testing"
	"webapp/pkg/repository/dbrepo"
)

var app Application
var expiredToken = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiYXVkIjoiZXhhbXBsZS5jb20iLCJleHAiOjE2ODUzODIyNzAsImlzcyI6ImV4YW1wbGUuY29tIiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMSJ9.ovBNmAOpfTSfkyRCMDjfA92mxXwN3w-lxGwhwBkD_XI`

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "very-secret"
	os.Exit(m.Run())
}
