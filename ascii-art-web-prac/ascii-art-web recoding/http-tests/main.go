package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/ok":
		http.Error(w, "successful", http.StatusOK)
		// w.WriteHeader(http.StatusOK)
	case "/badrequest":
		http.Error(w, "400 bad request", http.StatusBadRequest)
	case "/notfound":
		http.Error(w, "404 not found", http.StatusNotFound)
	case "/error":
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	default:
		http.Error(w, "Page not found \nAllowed routes:\n.../ok\n.../badrequest\n.../notfound\n.../error", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Program is live @port :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
