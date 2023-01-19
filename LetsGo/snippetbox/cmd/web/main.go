package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Dynamically Accepting addr flag via CLI
	addr := flag.String("addr", ":4000", "Http network Address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	// Serving Static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Binding loggers with server
	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Print("Server starting on http://localhost" + *addr)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
