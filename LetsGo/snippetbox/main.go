package main

import (
	"log"
	"net/http"
)

// A Request and Response Handler which would be converted into endpoint
// Route: localhost:4000/
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello From SnippetBox"))
}

// Route: localhost:4000/snippet/view
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a Snippet"))
}

// Route: localhost:4000/snippet/create
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a Snippet"))
}

func main() {
	// A new mux server using http.NewServeMux from http module/package
	mux := http.NewServeMux()

	// Request Handler using HandlerFunction followed by callback function
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	// Listening to Error While Serving and Throw Error Accordingly
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
