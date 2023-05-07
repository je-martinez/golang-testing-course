package main

import (
	"log"
	"net/http"
)

type Application struct{}

func main() {

	//App Config
	app := Application{}

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
