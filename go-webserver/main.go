package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request:", r.URL.Path)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request:", r.URL.Path)
		fmt.Fprintf(w, "Hi")
	})

	fmt.Println("Starting server on port 8080")
	defer fmt.Println("Server stopped!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
