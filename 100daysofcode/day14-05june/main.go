package main

// import (
// 	"fmt"
// 	"net/http"
// )

/*
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8090", nil)
}
*/

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func hello(w http.ResponseWriter, req *http.Request) {

// 	fmt.Fprintf(w, "hello\n")
// }

// func headers(w http.ResponseWriter, req *http.Request) {

// 	for name, headers := range req.Header {
// 		for _, h := range headers {
// 			fmt.Fprintf(w, "%v: %v\n", name, h)
// 		}
// 	}
// }

// func main() {

// 	http.HandleFunc("/hello", hello)
// 	http.HandleFunc("/headers", headers)
// 	fmt.Println("Running live at port 8090")
// 	http.ListenAndServe(":8090", nil)
// }
