package handler

import (
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) viewIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("resources/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
