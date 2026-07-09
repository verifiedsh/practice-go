package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path != "/" {
			http.Error(w, "page not found", http.StatusNotFound)
			return
		}
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	case "POST":
		text := r.FormValue("text")
		if text == "" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, text)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running live at port :8080")
	http.ListenAndServe(":8080", nil)
}
