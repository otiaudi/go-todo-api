package controller

import (
	"encoding/json"
	"net/http"

	"go-todo-api/model"
	"go-todo-api/views"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var data views.PostRequest
			json.NewDecoder(r.Body).Decode(&data)

			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				http.Error(w, "Some error occurred", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("✅ Todo created successfully"))

		case http.MethodGet:
			name := r.URL.Query().Get("name")
			data, err := model.ReadByName(name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)

		case http.MethodDelete:
			name := r.URL.Path[1:]
			if err := model.DeleteToDo(name); err != nil {
				http.Error(w, "Some error occurred", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(struct {
				Status string `json:"status"`
			}{"Todo deleted"})

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

// NEW: Handler for GET /all to retrieve all todos
// getAllTodos retrieves all todos from the database and returns them as JSON.

func getAllTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		todos, err := model.ReadAll()
		if err != nil {
			http.Error(w, "Error fetching todos: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todos)
	}
}
