package controller

import (
	"crud/model"
	"crud/view"
	"encoding/json"
	"net/http"
)

func crud() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			data := view.PostRequest{}
			json.NewDecoder(req.Body).Decode(&data)
			if err := model.CreateToDo(data.Name, data.ToDo); err != nil {
				resp.Write([]byte(err.Error()))
				return
			}
			resp.WriteHeader(http.StatusCreated)
			json.NewEncoder(resp).Encode(data)
		case http.MethodGet:
			name := req.URL.Query().Get("name")
			if len(name) > 0 {
				data, err := model.ReadByName(name)
				if err != nil {
					resp.Write([]byte(err.Error()))
				}
				resp.WriteHeader(http.StatusOK)
				json.NewEncoder(resp).Encode(data)
			} else {
				data, err := model.ReadAll()
				if err != nil {
					resp.Write([]byte(err.Error()))
				}
				resp.WriteHeader(http.StatusOK)
				json.NewEncoder(resp).Encode(data)
			}
		case http.MethodDelete:
			name := req.URL.Path[1:]
			if err := model.DeleteToDo(name); err != nil {
				resp.Write([]byte(err.Error()))
				return
			}
			resp.WriteHeader(http.StatusOK)
			json.NewEncoder(resp).Encode(struct {
				Status string `json:status`
			}{"Item deleted"})
		}
	}
}
