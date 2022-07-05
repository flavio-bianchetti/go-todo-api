package models

import "fmt"

func CreateTodo(name, description string) error {
	insertQ, err := con.Query(
		"INSERT INTO todo (name, description) VALUES(?, ?)", name, description,
  )
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

