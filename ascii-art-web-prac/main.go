package main

import (
	"ascii-art-web/helper"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helper.FormHandler)
	http.HandleFunc("/output", helper.ResultHandler)
	fmt.Println("Listening live at port :8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
