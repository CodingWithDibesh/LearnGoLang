package main

import (
	"crud/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)
	http.ListenAndServe("localhost:3000", mux)
}

func ping(resp http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	if req.Method == http.MethodGet {
		data := structs.Response{
			Code: http.StatusOK,
			Body: "pong",
		}
		json.NewEncoder(resp).Encode(data)
	}
}
