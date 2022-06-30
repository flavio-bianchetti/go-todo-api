package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/flavio-bianchetti/go-todo-api/views"
)

func ping() http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response {
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
