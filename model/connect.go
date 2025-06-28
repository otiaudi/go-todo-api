package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

var con *sql.DB

// Connect establishes a connection to the MySQL database and returns the connection object.

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:odero2020@tcp(localhost:3306)/GoDataBase")
	if err != nil {
		log.Fatal("Connection setup failed:", err)
	}
	fmt.Println("Connected to the database...")
	con = db
	// Check if the connection is actually established
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	fmt.Println("Database connection is alive.")
	con = db
	return db
}
