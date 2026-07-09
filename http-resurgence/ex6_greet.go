package main

import (
	"fmt"
	"net/http"
)

// ●     Register /v1/greet on apiMux — reads a name query parameter and
// responds with "Greetings, [name]!" or "Greetings, Stranger!" if name is missing.

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/v1/greet" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintln(w, "Greetings, Stranger!")
		return
	}
	fmt.Fprintf(w, "Greetings, %s!\n", name)
}
