package controller

import (
	"crud/model"
	"crud/view"
	"encoding/json"
	"net/http"
)

func create() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			data := view.PostRequest{}
			json.NewDecoder(req.Body).Decode(&data)
			if err := model.CreateToDo(data.Name, data.ToDo); err != nil {
				resp.Write([]byte(err.Error()))
				return
			}
			resp.WriteHeader(http.StatusCreated)
			json.NewEncoder(resp).Encode(data)
		}
	}
}
