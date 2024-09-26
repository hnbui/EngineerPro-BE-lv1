package main

import (
	"fmt"
	"log"
	"net/http"

	"ep-backend/controllers"
	"ep-backend/models"
)

func main() {
	// Initialize the database and handle any errors
	err := models.InitDB("users.db")
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}
	defer models.CloseDB() // Ensure that the database is closed when the program exits

	// Register HTTP handlers
	http.HandleFunc("/register", controllers.RegisterHandler)
	http.HandleFunc("/login", controllers.LoginHandler)

	// Start the server
	fmt.Println("Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
