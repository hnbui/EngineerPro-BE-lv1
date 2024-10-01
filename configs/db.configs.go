package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() error {
	db, err := sql.Open("mysql", "root:2301@tcp(127.0.0.1:3306)/engineerpro") // Connect to schema (database)
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MySQL database")

	return nil
}
