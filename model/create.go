package model

import "fmt"

func CreateTodo(name, todo string) error {
	insertQ, err := con.Query("INSERT INTO TodoList VALUES(?, ?)", name, todo)

	if err != nil {
		return err
	}
	defer insertQ.Close()
	return nil

}
func DeleteToDo(name string) error {
	insertQ, err := con.Query("DELETE FROM TodoList WHERE name=?", name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer insertQ.Close()
	return nil
}
