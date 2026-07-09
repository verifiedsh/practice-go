package main

import (
	"fmt"
	"net/http"
)

func main() {
	mainMux := http.NewServeMux()
	apiMux := http.NewServeMux()

	mainMux.HandleFunc("/method-inspector", methodHandler)
	mainMux.HandleFunc("/echo", echoHandler)
	mainMux.HandleFunc("/headers", headerHandler)
	mainMux.HandleFunc("/form", formHandler)
	mainMux.HandleFunc("/status", statusHandler)
	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))
	apiMux.HandleFunc("/v1/ping", pingHandler)
	apiMux.HandleFunc("/v1/greet", greetHandler)
	mainMux.HandleFunc("/render", renderHandler)
	fmt.Println("Ex 1-7 Live @:8080...")
	http.ListenAndServe(":8080", mainMux)
}
