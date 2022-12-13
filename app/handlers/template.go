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
		Name: "h_yamashita", // 自分の名前
		Age:  20,            // 自分の年齢
	}

	t, err := template.ParseFiles("templates/template.html")
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(w, d); err != nil {
		log.Fatal(err)
	}
}
