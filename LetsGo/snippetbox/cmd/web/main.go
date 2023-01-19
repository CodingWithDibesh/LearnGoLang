package main

import (
	"log"
	"net/http"
)

func main() {
	PORT := ":4000"
	mux := http.NewServeMux()

	// Serving Static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Server starting on http://localhost" + PORT)
	err := http.ListenAndServe(PORT, mux)
	log.Fatal(err)
}
