package main

import (
	"fmt"
	"net/http"
)

func redirectedHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path != "/v2" {
			http.Error(w, "page not found", http.StatusNotFound)
			return
		}
		fmt.Fprint(w, "Welcome to version 2")
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
