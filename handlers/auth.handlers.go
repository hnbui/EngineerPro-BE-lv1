package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ep-backend/users"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Allow only POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Request method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON request body into a User struct
	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Try to create the user in the database
	if err := users.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User %s is registered successfully", user.Username)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Allow only POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Request method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON request body into a Credential struct
	var credential users.Credential
	if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := users.AuthenticateUser(credential)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Logged in successfully")
}
