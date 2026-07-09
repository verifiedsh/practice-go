package helper

import (
	"html/template"
	"net/http"
)

func ResultHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		text := r.FormValue("text")
		banner := r.FormValue("banner")
		rows, er := LoadBanner(banner)
		if er != nil {
			http.Error(w, er.Error(), http.StatusInternalServerError)
			return
		}
		odds, er := ValidateInput(text)
		if odds != nil && er != nil {
			http.Error(w, er.Error(), http.StatusBadRequest)
			return
		}
		words := SplitInput(text)
		result := GenerateArt(words, rows)
		t, err := template.ParseFiles("templates/output.html")
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		t.ExecuteTemplate(w, "output.html", result)
	default:
		http.Error(w, "method not found", http.StatusMethodNotAllowed)
	}
}
