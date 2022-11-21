package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello World!"))

	d := struct {
		Name string
		Age  int
	}{
		Name: "西澤孟史",
		Age:  27,
	}

	t, err := template.ParseFiles("templates/hello.html")
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(w, d); err != nil {
		log.Fatal(err)
	}
}
