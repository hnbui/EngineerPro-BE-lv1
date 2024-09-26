package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ProfilePhotoURL string `json:"profile_photo_url"`
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *sql.DB

// HashPassword hashes the given password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

// CheckPasswordHash compares the provided password with the hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

// InitDB initializes the database connection and creates the necessary table if it doesn't exist.
func InitDB(dbName string) error {
	var err error
	db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		return fmt.Errorf("failed to initializing the database: %w", err) // Use fmt.Errorf to wrap the error with additional context
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL,
            profile_photo_url TEXT
        );`,
	)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

// Creates a new user in the database. The password is hashed before being stored.
func CreateUser(user User) error {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("can not hash password")
	}
	user.Password = string(hashedPassword)

	_, err = db.Exec("INSERT INTO users (username, password, profile_photo_url) VALUES (?, ?, ?)",
		user.Username, user.Password, user.ProfilePhotoURL)

	return err
}

// Authenticates a user by comparing the provided password with the hashed password stored in the database.
func AuthenticateUser(credential Credential) (User, error) {
	var user User

	row := db.QueryRow("SELECT username, password, profile_photo_url FROM users WHERE username = ?", credential.Username)
	err := row.Scan(&user.Username, &user.Password, &user.ProfilePhotoURL)

	if err != nil {
		// Return a proper error if the user is not found
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}

		return user, err
	}

	if !CheckPasswordHash(credential.Password, user.Password) {
		return user, fmt.Errorf("invalid credentials")
	}

	return user, nil
}

// Closes the database connection.
func CloseDB() {
	if db != nil {
		db.Close()
	}
}
