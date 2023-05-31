package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Application) routes() http.Handler {
	mux := chi.NewRouter()

	//Register Middlewar e
	mux.Use(middleware.Recoverer)
	mux.Use(app.addIpToContext)
	mux.Use(app.Session.LoadAndSave)

	//Register Router

	mux.Get("/", app.Home)
	mux.Post("/login", app.Login)

	//Static Assets
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
