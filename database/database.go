package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	// Establish connection to the database
	db, err := sql.Open("mysql", "root:2301@tcp(localhost:3306)/engineerpro")
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	// Check the database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	return db, nil
}
