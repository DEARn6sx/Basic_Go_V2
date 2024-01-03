package lesson

import (
	"fmt"
	"net/http"
)

func Ep3_netHttpHelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.",http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported",http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello World")
}


// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	lesson "github.com/DEAR/go-example/lesson/ep3"
// )

// func main() {
// 	http.HandleFunc("/hello", lesson.Ep3_netHttpHelloHandler)

// 	fmt.Printf("Starting server at port 8080\n")
// 	if err := http.ListenAndServe(":8080",nil); err != nil {
// 		log.Fatal(err)
// 	}
// }