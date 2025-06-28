package model

import "go-APIs/views"

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TodoList")
	if err != nil {
		return nil, err
	}

	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		// Scan the row into the data structure
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}

func ReadByName(name string) ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TodoList WHERE Name = ?", name)
	if err != nil {
		return nil, err
	}

	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		// Scan the row into the data structure
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}
