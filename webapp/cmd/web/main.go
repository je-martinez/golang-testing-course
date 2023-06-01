package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type Application struct {
	DSN     string
	DB      *sql.DB
	Session *scs.SessionManager
}

func main() {

	//App Config
	app := Application{}

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5433 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Postgres Connection")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = conn

	//get a session manager
	app.Session = getSession()

	//Application Routes
	mux := app.routes()

	//Print out a message

	log.Println("Starting server on port 8080....")

	//Start the server

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}
