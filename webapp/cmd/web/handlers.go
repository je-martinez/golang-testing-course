package main

import (
	"html/template"
	"net/http"
	"path"
)

var pathToTemplates = "./templates/"

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home.page.gohtml", &TemplateData{})
}

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *Application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {

	//Parse Template from Disk
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t))

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
