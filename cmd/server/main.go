package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerfunc(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to my awesoem site")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", handlerfunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
