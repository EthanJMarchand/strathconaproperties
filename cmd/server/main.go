package main

import (
	"fmt"
	"log"
	"net/http"
)

type homeHandler struct {
}

func newHomeHandler() *homeHandler {
	return &homeHandler{}
}

func (h homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to my awesoem site")
	w.Header().Set("Content-Type", "text/html")
	if err != nil {
		return
	}
}

type aboutHandler struct{}

func newAboutHandler() *aboutHandler {
	return &aboutHandler{}
}

func (h aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	thing := r.PathValue("thing")

	_, err := fmt.Fprintf(w, "Welcome to the about page. The thing is: %v ", thing)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.Handle("GET /", newHomeHandler())
	http.Handle("GET /about/{thing}", newAboutHandler())
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
