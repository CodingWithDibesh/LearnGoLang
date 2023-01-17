package main

import (
	"crud/controller"
	"fmt"
	"net/http"
)

func main() {
	server := "localhost:3000"
	fmt.Println("Server started at http://", server)
	http.ListenAndServe(server, controller.Register())
}
