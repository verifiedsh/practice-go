// Exercise 7: Simple Redirector (Status Codes)
// Goal: Create a /legacy route. Whenever a user hits this endpoint, permanently redirect them to a new route /v2 with a friendly "Welcome to version 2" message.
// Key Tasks:
// Redirect traffic using the http.Redirect helper function.
// Use the proper status code for a permanent move (http.StatusMovedPermanently).
package main

import (
	"net/http"
)

func redirectorHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path != "/legacy" {
			http.Error(w, "page not found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, "http://localhost:8080/v2", http.StatusMovedPermanently)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
