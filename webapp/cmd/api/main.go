package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"webapp/pkg/repository"
	"webapp/pkg/repository/dbrepo"
)

const port = 8090

type Application struct {
	DSN       string
	DB        repository.DatabaseRepo
	Domain    string
	JWTSecret string
}

func main() {
	var app Application

	flag.StringVar(&app.Domain, "domain", "example.com", "Domain for application, e.g. company.com")
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5433 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Postgres Connection")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "very-secret", "signing secret")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	log.Printf("Starting api on port %d", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
}
