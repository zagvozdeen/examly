package main

import (
	"html/template"
	"net/http"
)

func (app *application) viewIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("resources/views/index.html")
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		app.log.Err(err).Msg("failed to execute template")
	}
}
