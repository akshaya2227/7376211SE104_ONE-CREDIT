package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Akshaya U S")
	})
	http.HandleFunc("/regnum", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376211SE104")
	})
	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Blue")
	})

	fmt.Printf("Server running (port=8081), route: http://localhost:8081/\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
