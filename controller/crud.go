package controller

import (
	"encoding/json"
	// "fmt"
	"go-APIs/model"
	"go-APIs/views"
	"net/http"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			// fmt.Println(data)

			//call model function to create todo
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some error occured"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Todo created successfully"))

		} else if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			data, err := model.ReadByName(name)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[1:]

			if err := model.DeleteToDo(name); err != nil {
				w.Write([]byte("Some errors"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Status string `json:"status"`
			}{"item deleted"})
		}
	}
}
