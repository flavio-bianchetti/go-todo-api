package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/flavio-bianchetti/go-todo-api/models"
	"github.com/flavio-bianchetti/go-todo-api/views"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			if err := models.CreateTodo(data.Name, data.Description); err != nil {
				w.Write([]byte("Create method failed: " + err.Error()))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			if name != "" {
				data, err := models.ReadByName(name)
				if err != nil {
					w.Write([]byte(err.Error()))
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(data)
			} else {
				data, err := models.ReadAll()
				if err != nil {
					w.Write([]byte(err.Error()))
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(data)
			}
		} else if r.Method == http.MethodDelete {
			id := r.URL.Path[1:]
			if err := models.DeleteTodo(id); err != nil {
				w.Write([]byte("Delete method failed: " + err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(nil)
		}
	}
}
