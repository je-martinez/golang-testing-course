package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

var pathToTemplates = "./templates/"

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	var td = make(map[string]any)

	if app.Session.Exists(r.Context(), "test") {
		msg := app.Session.GetString(r.Context(), "test")
		td["test"] = msg
	} else {
		app.Session.Put(r.Context(), "test", "Hit this page at "+time.Now().UTC().String())
	}

	_ = app.render(w, r, "home.page.gohtml", &TemplateData{Data: td})
}

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *Application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {

	//Parse Template from Disk
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "base.layout.gohtml"))

	data.IP = app.ipFromContext(r.Context())

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return err
	}

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return nil
	}

	return nil
}

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	form := NewForm(r.PostForm)
	form.Required("email", "password")

	if !form.Valid() {
		fmt.Fprint(w, "Failed validation")
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	log.Println(email, password)
	fmt.Fprintf(w, email)

}
