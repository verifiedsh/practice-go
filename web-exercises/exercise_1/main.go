package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template
var err error

type values struct {
	Text   string
	Result string
	ErrMsg string
}

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/echo", EchoHandler)
	fmt.Println("Live...")
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := values{
		ErrMsg: "",
	}
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		data.ErrMsg = "page not found"
		tmpl.ExecuteTemplate(w, "404.html", data)
		return
	}

	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data.ErrMsg = "Internal server error"
		tmpl.ExecuteTemplate(w, "500.html", nil)
		return
	}
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	data := values{
		Text:   "",
		Result: "",
		ErrMsg: "",
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		data.ErrMsg = "page not found"
		tmpl.ExecuteTemplate(w, "404.html", nil)
		return
	}
	text := r.FormValue("text")
	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		data.ErrMsg = "page not found"
		tmpl.ExecuteTemplate(w, "400.html", nil)
		return
	}
	data.Result = `Weldone, You typed in "` + text + `".`

	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data.ErrMsg = "internal server error"
		tmpl.ExecuteTemplate(w, "500.html", nil)
		return
	}
}
