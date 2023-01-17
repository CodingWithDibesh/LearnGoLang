package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, rep *http.Request) {
		resp.Write([]byte("Hello World"))
	})
	http.ListenAndServe("localhost:3000", mux)
}
