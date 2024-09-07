package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":8080"

func Root(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Root")
	if err != nil {
		return
	}
}

func main() {

	http.HandleFunc("/", Root)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
