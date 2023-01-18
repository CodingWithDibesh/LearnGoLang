package main

import (
	"log"
	"net/http"
)

// A Request and Response Handler which would be converted into endpoint
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From SnippetBox"))
}

func main() {
	// A new mux server using http.NewServeMux from http module/package
	mux := http.NewServeMux()

	// Request Handler using HandlerFunction followed by callback function
	mux.HandleFunc("/", home)

	log.Print("Starting server on :4000")
	// Listening to Error While Serving and Throw Error Accordingly
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
