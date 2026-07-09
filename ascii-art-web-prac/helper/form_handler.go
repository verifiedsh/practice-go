package helper

import (
	"html/template"
	"net/http"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path != "/" {
			http.Error(w, "page not found", http.StatusNotFound)
			return
		}
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		t.Execute(w, nil)
	default:
		http.Error(w, "method not found", http.StatusMethodNotAllowed)
	}
}

