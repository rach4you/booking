package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders templates using html / template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("../../templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

// RenderTemplate renders templates using html / template
/*
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	t, err := template.ParseFiles("../../templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
