package main

import (
	"fmt"
	"net/http"
	"go-APIs/controller" // 👈 This matches your module name in go.mod
	"go-APIs/model"
	"log"
)

func main() {
	// Create multiplexer
	mux := controller.Register() // 👈 Call the Register function from controller package
	db := model.Connect()
	defer db.Close()
	fmt.Println("Server running at http://localhost:3002/ping")
	log.Fatal(http.ListenAndServe(":3002", mux))
}
