package controller

import (
	"crud/view"
	"encoding/json"
	"net/http"
)

func ping() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			data := view.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			resp.WriteHeader(http.StatusOK)
			json.NewEncoder(resp).Encode(data)
		}
	}
}
