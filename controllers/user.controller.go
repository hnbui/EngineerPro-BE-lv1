package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ep-backend/models"
)

// RegisterHandler handles user registration requests
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Allow only POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Request method not allowed", http.StatusMethodNotAllowed)

		return
	}

	// Decode the JSON request body into a User struct
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)

		return
	}

	// Try to create the user in the database
	if err := models.CreateUser(user); err != nil {
		http.Error(w, "User already exists or cannot create user", http.StatusConflict)

		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User %s is registered successfully", user.Username)
}

// LoginHandler handles user login requests
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Allow only POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Request method not allowed", http.StatusMethodNotAllowed)

		return
	}

	// Decode the JSON request body into a Credential struct
	var credential models.Credential
	if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)

		return
	}

	// Authenticate the user
	_, err := models.AuthenticateUser(credential)
	if err != nil {
		// Handle the case where the user is not found
		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)

			return
		}
		// Handle other errors (e.g., database errors)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)

		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Logged in successfully")
}
