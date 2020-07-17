package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		fmt.Print(err)
		tmpl.Execute(w, "context")
	})

	http.ListenAndServe(":8080", nil)
}
