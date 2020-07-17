package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is home page")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is about page")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is contact page")
	})

	http.ListenAndServe(":8080", nil)
}
