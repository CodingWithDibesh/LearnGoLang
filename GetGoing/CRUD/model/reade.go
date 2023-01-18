package model

import "crud/view"

func ReadAll() ([]view.PostRequest, error) {
	rows, err := connection.Query("SELECT * FROM TODO;")
	if err != nil {
		return nil, err
	}
	todos := []view.PostRequest{}
	for rows.Next() {
		data := view.PostRequest{}
		rows.Scan(&data.Name, &data.ToDo)
		todos = append(todos, data)
	}
	return todos, nil
}

func ReadByName(name string) ([]view.PostRequest, error) {
	rows, err := connection.Query("SELECT * FROM TODO WHERE name=?;", name)
	if err != nil {
		return nil, err
	}
	todos := []view.PostRequest{}
	for rows.Next() {
		data := view.PostRequest{}
		rows.Scan(&data.Name, &data.ToDo)
		todos = append(todos, data)
	}
	return todos, nil
}
