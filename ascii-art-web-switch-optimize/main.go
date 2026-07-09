package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type PageData struct {
	Text   string
	Banner string
	Result string
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ascii-art", PostHandler)
	http.HandleFunc("/ascii-art-switch", SwitchHandler)
	fmt.Println("Live...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		page := PageData{
			Text:   "",
			Banner: "",
			Result: "",
		}
		err := tmpl.Execute(w, page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		text := r.FormValue("text")
		banner := r.FormValue("banner")
		bannerFile := banner + ".txt"
		rows, err := LoadBanner(bannerFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		odd, err := Validateinput(text)
		if odd != 0 && err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		words := SplitInput(text)
		result := GenerateArt(rows, words)
		var data = PageData{
			Text:   text,
			Banner: banner,
			Result: result,
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func SwitchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		text := r.URL.Query().Get("text")
		banner := r.URL.Query().Get("banner")
		bannerFile := banner + ".txt"
		rows, err := LoadBanner(bannerFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		odd, err := Validateinput(text)
		if odd != 0 && err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		words := SplitInput(text)
		result := GenerateArt(rows, words)
		var data = PageData{
			Text:   text,
			Banner: banner,
			Result: result,
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
