package main

import (
	"log"
	"net/http"

	"github.com/mow-nishizawa/go_sample/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head>
					<title>サンプル</title>
				</head>
				<body>Hello World!</body>
			</html>
		`))
	})
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/template", handlers.TemplateHandler)
	http.HandleFunc("/request_parse", handlers.RequestParseHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
