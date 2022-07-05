package models

import (
	"github.com/flavio-bianchetti/go-todo-api/views"
)

func ReadAll() ([]views.GetRequest, error) {
	rows, err := con.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}
	todos := []views.GetRequest{}
	for rows.Next() {
		data := views.GetRequest{}
		rows.Scan(&data.Id, &data.Name, &data.Description)
		todos = append(todos, data)
	}
	return todos, nil
}

func ReadByName(name string) ([]views.GetRequest, error) {
	rows, err := con.Query("SELECT * FROM todo WHERE name=?", name)
	if err != nil {
		return nil, err
	}
	todos := []views.GetRequest{}
	for rows.Next() {
		data := views.GetRequest{}
		rows.Scan(&data.Id, &data.Name, &data.Description)
		todos = append(todos, data)
	}
	return todos, nil
}
