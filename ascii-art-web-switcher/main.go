package main

import(
	"fmt"
	"net/http"
	"strings"
	"html/template"
)

type pagedata struct {
	Text string
	Banner string
	Result string
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/ascii-art", handleAsciiArt)
	http.HandleFunc("/ascii-art-switch", handleSwitch)
	fmt.Println("Live...")
	http.ListenAndServe(":8080", nil)
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}	
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	if text == "" || banner == "" {
		http.Error(w, "empty parameters, provide valid entries", http.StatusBadRequest)
		return
	}
	if banner != "shadow" && banner != "standard" && banner != "thinkertoy" {
		http.Error(w, "invalid banner", http.StatusBadRequest)
		return
	}
	text = strings.ReplaceAll(text, "\r\n", "\n")
	words := strings.Split(text, "\n")
	rows, err := LoadBanner(banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result := GenerateArt(words, rows)
	page := pagedata {
		Text: text,
		Banner: banner,
		Result: result,
	}
	err = tmpl.Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleSwitch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}	
	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")
	if text == "" || banner == "" {
		http.Error(w, "empty parameters, provide valid entries", http.StatusBadRequest)
		return
	}
	if banner != "shadow" && banner != "standard" && banner != "thinkertoy" {
		http.Error(w, "invalid banner", http.StatusBadRequest)
		return
	}
	text = strings.ReplaceAll(text, "\r\n", "\n")
	words := strings.Split(text, "\n")
	rows, err := LoadBanner(banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result := GenerateArt(words, rows)
	page := pagedata {
		Text: text,
		Banner: banner,
		Result: result,
	}
	err = tmpl.Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}