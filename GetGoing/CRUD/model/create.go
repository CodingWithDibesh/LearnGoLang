package model

import "fmt"

func CreateToDo(name, todo string) error {
	insertQ, err := connection.Query("INSERT INTO TODO VALUES(?,?)", name, todo)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer insertQ.Close()
	return nil
}

func DeleteToDo(name string) error {
	deleteQ, err := connection.Query("DELETE FROM TODO WHERE name=?", name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer deleteQ.Close()
	return nil
}
