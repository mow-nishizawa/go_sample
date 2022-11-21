package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	d := struct {
		Name string
		Age  int
	}{
		Name: "西澤孟史",
		Age:  27,
	}

	t, err := template.ParseFiles("templates/template.html")
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(w, d); err != nil {
		log.Fatal(err)
	}
}
