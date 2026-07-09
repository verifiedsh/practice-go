// Exercise 1: Basic Ping-Pong Server
// Goal: Build a minimal web server that listens on port 8080 and responds with "pong" when a user visits the /ping route.
// Tasks:
// Create a route handler for /ping using http.HandleFunc.
// Use w.Write() or fmt.Fprint() to send a plain text response.
// Start the server on port :8080 using http.ListenAndServe.
package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path != "/ping" {
			http.Error(w, "page not found", http.StatusNotFound)
			return
		}
		// fmt.Fprint(w,"pong")
		w.Write([]byte("pong"))
	default:
		http.Error(w, "method not allowed", http.StatusNotFound)
	}
}
