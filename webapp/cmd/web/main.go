package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type Application struct {
	Session *scs.SessionManager
}

func main() {

	//App Config
	app := Application{}

	//get a session manager
	app.Session = getSession()

	//Application Routes
	mux := app.routes()

	//Print out a message

	log.Println("Starting server on port 8080....")

	//Start the server

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}
