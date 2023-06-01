package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
	"webapp/pkg/data"
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
	IP    string
	Data  map[string]any
	Error string
	Flash string
	User  data.User
}

func (app *Application) render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) error {

	//Parse Template from Disk
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "base.layout.gohtml"))

	td.IP = app.ipFromContext(r.Context())

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return err
	}

	td.Error = app.Session.PopString(r.Context(), "error")
	td.Flash = app.Session.PopString(r.Context(), "flash")

	err = parsedTemplate.Execute(w, td)
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
		//redirect to login page with error message
		app.Session.Put(r.Context(), "error", "Invalid login credentials")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := app.DB.GetUserByEmail(email)

	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid login!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	log.Println(user)
	log.Println(password)

	//authenticate the user

	//if not authenticated then redirect with error

	//prevent fixation attack
	_ = app.Session.RenewToken(r.Context())

	//store success message in session

	//redirect to some other page
	app.Session.Put(r.Context(), "flash", "Successfully logged in!")
	http.Redirect(w, r, "/user/profile", http.StatusSeeOther)
}

func (app *Application) Profile(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "profile.page.gohtml", &TemplateData{})
}
