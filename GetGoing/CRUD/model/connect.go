package model

import (
	"database/sql"
	"fmt"
	"log"
)

var connection *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@/mysql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	connection = db
	return db
}
