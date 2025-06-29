package main

import (
	"fmt"
	"log"
	"net/http"

	"go-todo-api/controller"
	"go-todo-api/model"
)

func main() {
	mux := controller.Register()

	db := model.Connect()
	defer db.Close()

	fmt.Println("Server running at http://localhost:3002/ping")
	log.Fatal(http.ListenAndServe(":3002", mux))
}
