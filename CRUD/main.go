package main

import (
	"crud/controller"
	"crud/model"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server := "localhost:3000"
	db := model.Connect()
	defer db.Close()
	fmt.Println("Server started at http://", server)
	log.Fatal(http.ListenAndServe(server, controller.Register()))
}
