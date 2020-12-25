package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	log.Println("Running on port", port)
	http.ListenAndServe(port, nil)
}
